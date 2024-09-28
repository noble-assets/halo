// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/header"
	"cosmossdk.io/core/store"
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v2/types"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/types/entitlements"
)

type Keeper struct {
	Denom      string
	Underlying string

	Schema        collections.Schema
	storeService  store.KVStoreService
	eventService  event.Service
	headerService header.Service

	Owner  collections.Item[string]
	Nonces collections.Map[[]byte, uint64]

	AggregatorOwner collections.Item[string]
	LastRoundId     collections.Sequence
	NextPrice       collections.Item[math.Int]
	Rounds          collections.Map[uint64, aggregator.RoundData]

	EntitlementsOwner  collections.Item[string]
	Paused             collections.Item[bool]
	PublicCapabilities collections.Map[string, bool]
	RoleCapabilities   collections.Map[collections.Pair[string, uint64], bool]
	UserRoles          collections.Map[collections.Pair[[]byte, uint64], bool]

	addressCodec      address.Codec
	accountKeeper     types.AccountKeeper
	bankKeeper        types.BankKeeper
	interfaceRegistry codectypes.InterfaceRegistry
}

func NewKeeper(
	cdc codec.Codec,
	storeService store.KVStoreService,
	eventService event.Service,
	headerService header.Service,
	denom string,
	underlying string,
	addressCodec address.Codec,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	interfaceRegistry codectypes.InterfaceRegistry,
) *Keeper {
	builder := collections.NewSchemaBuilder(storeService)

	keeper := &Keeper{
		Denom:      denom,
		Underlying: underlying,

		addressCodec:  addressCodec,
		storeService:  storeService,
		eventService:  eventService,
		headerService: headerService,

		Owner:  collections.NewItem(builder, types.OwnerKey, "owner", collections.StringValue),
		Nonces: collections.NewMap(builder, types.NoncePrefix, "nonces", collections.BytesKey, collections.Uint64Value),

		AggregatorOwner: collections.NewItem(builder, aggregator.OwnerKey, "aggregator_owner", collections.StringValue),
		LastRoundId:     collections.NewSequence(builder, aggregator.LastRoundIDKey, "aggregator_last_round_id"),
		NextPrice:       collections.NewItem(builder, aggregator.NextPriceKey, "aggregator_next_price", sdk.IntValue),
		Rounds:          collections.NewMap(builder, aggregator.RoundPrefix, "aggregator_rounds", collections.Uint64Key, codec.CollValue[aggregator.RoundData](cdc)),

		EntitlementsOwner:  collections.NewItem(builder, entitlements.OwnerKey, "entitlements_owner", collections.StringValue),
		Paused:             collections.NewItem(builder, entitlements.PausedKey, "entitlements_paused", collections.BoolValue),
		PublicCapabilities: collections.NewMap(builder, entitlements.PublicPrefix, "entitlements_public_capabilities", collections.StringKey, collections.BoolValue),
		RoleCapabilities:   collections.NewMap(builder, entitlements.CapabilityPrefix, "entitlements_role_capabilities", collections.PairKeyCodec(collections.StringKey, collections.Uint64Key), collections.BoolValue),
		UserRoles:          collections.NewMap(builder, entitlements.UserPrefix, "entitlements_user_roles", collections.PairKeyCodec(collections.BytesKey, collections.Uint64Key), collections.BoolValue),

		accountKeeper:     accountKeeper,
		bankKeeper:        bankKeeper,
		interfaceRegistry: interfaceRegistry,
	}

	schema, err := builder.Build()
	if err != nil {
		panic(err)
	}

	keeper.Schema = schema
	return keeper
}

// SetBankKeeper overwrites the bank keeper used in this module.
func (k *Keeper) SetBankKeeper(bankKeeper types.BankKeeper) {
	k.bankKeeper = bankKeeper
}

// SendRestrictionFn executes necessary checks against all USYC transfers.
func (k *Keeper) SendRestrictionFn(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) (newToAddr sdk.AccAddress, err error) {
	if amount := amt.AmountOf(k.Denom); !amount.IsZero() {
		burning := !fromAddr.Equals(types.ModuleAddress) && toAddr.Equals(types.ModuleAddress)
		minting := fromAddr.Equals(types.ModuleAddress) && !toAddr.Equals(types.ModuleAddress)

		if !minting {
			if !k.CanCall(ctx, fromAddr, "transfer") {
				return toAddr, fmt.Errorf("%s cannot transfer %s", fromAddr.String(), k.Denom)
			}
		}

		if !burning {
			if !k.CanCall(ctx, toAddr, "transfer") {
				return toAddr, fmt.Errorf("%s cannot transfer %s", toAddr.String(), k.Denom)
			}
		}
	}

	return toAddr, nil
}

// VerifyWithdrawSignature ensures that the owner has signed a withdrawal.
func (k *Keeper) VerifyWithdrawSignature(ctx context.Context, recipient sdk.AccAddress, amount math.Int, signature []byte) bool {
	owner, _ := k.addressCodec.StringToBytes(k.GetOwner(ctx))
	account := k.accountKeeper.GetAccount(ctx, owner)

	if account == nil || account.GetPubKey() == nil {
		return false
	}

	nonce, err := k.IncrementNonce(ctx, recipient)
	if err != nil {
		return false
	}

	bz, err := json.Marshal(types.WithdrawSignatureWrapper{
		Data: types.WithdrawSignatureData{
			Recipient: recipient,
			Amount:    amount,
			Nonce:     nonce,
		},
	})
	if err != nil {
		return false
	}

	return account.GetPubKey().VerifySignature(bz, signature)
}

// burnCoins is an internal helper function to burn.
func (k *Keeper) burnCoins(ctx context.Context, sender sdk.AccAddress, coins sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins)
	if err != nil {
		return errors.Wrap(err, "unable to transfer from account to module")
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return errors.Wrap(err, "unable to burn from module")
	}

	return nil
}

// mintCoins is an internal helper function to mint.
func (k *Keeper) mintCoins(ctx context.Context, recipient sdk.AccAddress, coins sdk.Coins) error {
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return errors.Wrap(err, "unable to mint to module")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)
	if err != nil {
		return errors.Wrap(err, "unable to transfer from module to account")
	}

	return nil
}

// depositFor is an internal helper function to deposit.
func (k *Keeper) depositFor(ctx context.Context, signer sdk.AccAddress, recipient sdk.AccAddress, underlying math.Int) (amount math.Int, err error) {
	lastRoundId := k.GetLastRoundId(ctx)
	round, found := k.GetRound(ctx, lastRoundId)
	if !found {
		return math.Int{}, fmt.Errorf("round %d not found", lastRoundId)
	}
	amount = underlying.QuoRaw(10000).MulRaw(10000)
	amount = amount.MulRaw(100000000).Quo(round.Answer)

	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	err = k.mintCoins(ctx, recipient, coins)
	if err != nil {
		return amount, err
	}

	coins = sdk.NewCoins(sdk.NewCoin(k.Underlying, underlying))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, coins)
	if err != nil {
		return amount, errors.Wrap(err, "unable to transfer from account to module")
	}
	return amount, k.eventService.EventManager(ctx).Emit(ctx, &types.Deposit{
		From:   signer.String(),
		Amount: underlying,
	})
}

// withdrawTo is an internal helper function to withdraw.
func (k *Keeper) withdrawTo(ctx context.Context, signer sdk.AccAddress, recipient sdk.AccAddress, amount math.Int) (underlying math.Int, err error) {
	lastRoundId := k.GetLastRoundId(ctx)
	round, found := k.GetRound(ctx, lastRoundId)
	if !found {
		return math.Int{}, fmt.Errorf("round %d not found", lastRoundId)
	}
	underlying = amount.Mul(round.Answer).QuoRaw(100000000)
	underlying = underlying.QuoRaw(10000).MulRaw(10000)

	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, amount))
	err = k.burnCoins(ctx, signer, coins)
	if err != nil {
		return underlying, err
	}

	coins = sdk.NewCoins(sdk.NewCoin(k.Underlying, underlying))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)
	if err != nil {
		return underlying, errors.Wrap(err, "unable to transfer from module to account")
	}

	return underlying, k.eventService.EventManager(ctx).Emit(ctx, &types.Withdrawal{
		To:     recipient.String(),
		Amount: underlying,
	})
}
