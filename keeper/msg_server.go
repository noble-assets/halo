package keeper

import (
	"bytes"
	"context"
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v2/types"
	"github.com/noble-assets/halo/v2/types/entitlements"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Deposit(ctx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	method := sdk.MsgTypeURL(msg)

	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	amount, err := k.depositFor(ctx, signer, signer, msg.Amount)
	return &types.MsgDepositResponse{Amount: amount}, err
}

func (k msgServer) DepositFor(ctx context.Context, msg *types.MsgDepositFor) (*types.MsgDepositResponse, error) {
	method := sdk.MsgTypeURL(msg)

	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	recipient, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}
	if !bytes.Equal(signer, recipient) {
		if !k.CanCall(ctx, recipient, "transfer") {
			return nil, fmt.Errorf("%s cannot receive %s", msg.Recipient, k.Denom)
		}
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	amount, err := k.depositFor(ctx, signer, recipient, msg.Amount)
	return &types.MsgDepositResponse{Amount: amount}, err
}

func (k msgServer) Withdraw(ctx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	method := sdk.MsgTypeURL(msg)

	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	if !k.VerifyWithdrawSignature(ctx, signer, msg.Amount, msg.Signature) {
		return nil, types.ErrInvalidSignature
	}

	underlying, err := k.withdrawTo(ctx, signer, signer, msg.Amount)
	return &types.MsgWithdrawResponse{Amount: underlying}, err
}

func (k msgServer) WithdrawTo(ctx context.Context, msg *types.MsgWithdrawTo) (*types.MsgWithdrawResponse, error) {
	method := sdk.MsgTypeURL(msg)

	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	recipient, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}
	if !bytes.Equal(signer, recipient) {
		if !k.CanCall(ctx, recipient, "transfer") {
			return nil, fmt.Errorf("%s cannot receive %s", msg.Recipient, k.Denom)
		}
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	if !k.VerifyWithdrawSignature(ctx, recipient, msg.Amount, msg.Signature) {
		return nil, types.ErrInvalidSignature
	}

	underlying, err := k.withdrawTo(ctx, signer, recipient, msg.Amount)
	return &types.MsgWithdrawResponse{Amount: underlying}, err
}

func (k msgServer) WithdrawToAdmin(ctx context.Context, msg *types.MsgWithdrawToAdmin) (*types.MsgWithdrawResponse, error) {
	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	from, err := k.addressCodec.StringToBytes(msg.From)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode from address %s", msg.From)
	}
	recipient, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}

	underlying, err := k.withdrawTo(ctx, from, recipient, msg.Amount)
	return &types.MsgWithdrawResponse{Amount: underlying}, err
}

func (k msgServer) Burn(ctx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	method := sdk.MsgTypeURL(msg)

	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}
	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, msg.Amount))
	err = k.burnCoins(ctx, signer, coins)

	return &types.MsgBurnResponse{}, err
}

func (k msgServer) BurnFor(ctx context.Context, msg *types.MsgBurnFor) (*types.MsgBurnForResponse, error) {
	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	from, err := k.addressCodec.StringToBytes(msg.From)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode from address %s", msg.From)
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, msg.Amount))
	err = k.burnCoins(ctx, from, coins)

	return &types.MsgBurnForResponse{}, err
}

func (k msgServer) Mint(ctx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	to, err := k.addressCodec.StringToBytes(msg.To)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode to address %s", msg.To)
	}
	if !k.CanCall(ctx, to, "transfer") {
		return nil, fmt.Errorf("%s cannot transfer %s", msg.To, k.Denom)
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, msg.Amount))
	err = k.mintCoins(ctx, to, coins)

	return &types.MsgMintResponse{}, err
}

func (k msgServer) TradeToFiat(ctx context.Context, msg *types.MsgTradeToFiat) (*types.MsgTradeToFiatResponse, error) {
	signer, err := k.addressCodec.StringToBytes(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	recipient, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}
	if !k.HasRole(ctx, recipient, entitlements.ROLE_LIQUIDITY_PROVIDER) {
		return nil, types.ErrInvalidLiquidityProvider
	}
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid amount %s", msg.Amount.String())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, recipient,
		sdk.NewCoins(sdk.NewCoin(k.Underlying, msg.Amount)),
	)

	return &types.MsgTradeToFiatResponse{}, err
}

func (k msgServer) TransferOwnership(ctx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	owner := k.GetOwner(ctx)
	if owner == "" {
		return nil, types.ErrNoOwner
	}
	if msg.Signer != owner {
		return nil, errors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	if msg.NewOwner == owner {
		return nil, types.ErrSameOwner
	}

	if err := k.SetOwner(ctx, msg.NewOwner); err != nil {
		return nil, err
	}

	return &types.MsgTransferOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}
