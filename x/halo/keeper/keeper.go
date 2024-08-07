package keeper

import (
	"encoding/json"
	"fmt"

	sdkerrors "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types"
)

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey

	Denom      string
	Underlying string

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	denom string,
	underlying string,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		Denom:      denom,
		Underlying: underlying,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// SetBankKeeper overwrites the bank keeper used in this module.
func (k *Keeper) SetBankKeeper(bankKeeper types.BankKeeper) {
	k.bankKeeper = bankKeeper
}

// SendRestrictionFn executes necessary checks against all USYC transfers.
func (k *Keeper) SendRestrictionFn(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) (newToAddr sdk.AccAddress, err error) {
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
func (k *Keeper) VerifyWithdrawSignature(ctx sdk.Context, recipient sdk.AccAddress, amount sdk.Int, signature []byte) bool {
	owner := sdk.MustAccAddressFromBech32(k.GetOwner(ctx))
	account := k.accountKeeper.GetAccount(ctx, owner)

	if account == nil || account.GetPubKey() == nil {
		return false
	}

	bz, err := json.Marshal(types.WithdrawSignatureWrapper{
		Data: types.WithdrawSignatureData{
			Recipient: recipient,
			Amount:    amount,
			Nonce:     k.IncrementNonce(ctx, recipient),
		},
	})
	if err != nil {
		return false
	}

	return account.GetPubKey().VerifySignature(bz, signature)
}

// burnCoins is an internal helper function to burn.
func (k *Keeper) burnCoins(ctx sdk.Context, sender sdk.AccAddress, coins sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins)
	if err != nil {
		return sdkerrors.Wrap(err, "unable to transfer from account to module")
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return sdkerrors.Wrap(err, "unable to burn from module")
	}

	return nil
}

// mintCoins is an internal helper function to mint.
func (k *Keeper) mintCoins(ctx sdk.Context, recipient sdk.AccAddress, coins sdk.Coins) error {
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return sdkerrors.Wrap(err, "unable to mint to module")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)
	if err != nil {
		return sdkerrors.Wrap(err, "unable to transfer from module to account")
	}

	return nil
}

// depositFor is an internal helper function to deposit.
func (k *Keeper) depositFor(ctx sdk.Context, signer sdk.AccAddress, recipient sdk.AccAddress, underlying sdk.Int) (amount sdk.Int, err error) {
	lastRoundId := k.GetLastRoundId(ctx)
	round, found := k.GetRound(ctx, lastRoundId)
	if !found {
		return sdk.Int{}, fmt.Errorf("round %d not found", lastRoundId)
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
		return amount, sdkerrors.Wrap(err, "unable to transfer from account to module")
	}

	return amount, ctx.EventManager().EmitTypedEvent(&types.Deposit{
		From:   signer.String(),
		Amount: underlying,
	})
}

// withdrawTo is an internal helper function to withdraw.
func (k *Keeper) withdrawTo(ctx sdk.Context, signer sdk.AccAddress, recipient sdk.AccAddress, amount sdk.Int) (underlying sdk.Int, err error) {
	lastRoundId := k.GetLastRoundId(ctx)
	round, found := k.GetRound(ctx, lastRoundId)
	if !found {
		return sdk.Int{}, fmt.Errorf("round %d not found", lastRoundId)
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
		return underlying, sdkerrors.Wrap(err, "unable to transfer from module to account")
	}

	return underlying, ctx.EventManager().EmitTypedEvent(&types.Withdrawal{
		To:     recipient.String(),
		Amount: underlying,
	})
}
