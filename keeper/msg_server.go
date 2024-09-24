package keeper

import (
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

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	method := sdk.MsgTypeURL(msg)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
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

func (k msgServer) DepositFor(goCtx context.Context, msg *types.MsgDepositFor) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	method := sdk.MsgTypeURL(msg)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}
	if !signer.Equals(recipient) {
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

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	method := sdk.MsgTypeURL(msg)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
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

func (k msgServer) WithdrawTo(goCtx context.Context, msg *types.MsgWithdrawTo) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	method := sdk.MsgTypeURL(msg)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.CanCall(ctx, signer, method) {
		return nil, fmt.Errorf("%s cannot execute %s", msg.Signer, method)
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}
	if !signer.Equals(recipient) {
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

func (k msgServer) WithdrawToAdmin(goCtx context.Context, msg *types.MsgWithdrawToAdmin) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode from address %s", msg.From)
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode recipient address %s", msg.Recipient)
	}

	underlying, err := k.withdrawTo(ctx, from, recipient, msg.Amount)
	return &types.MsgWithdrawResponse{Amount: underlying}, err
}

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	method := sdk.MsgTypeURL(msg)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
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

func (k msgServer) BurnFor(goCtx context.Context, msg *types.MsgBurnFor) (*types.MsgBurnForResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	from, err := sdk.AccAddressFromBech32(msg.From)
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

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	to, err := sdk.AccAddressFromBech32(msg.To)
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

func (k msgServer) TradeToFiat(goCtx context.Context, msg *types.MsgTradeToFiat) (*types.MsgTradeToFiatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode signer address %s", msg.Signer)
	}
	if !k.HasRole(ctx, signer, entitlements.ROLE_FUND_ADMIN) {
		return nil, types.ErrInvalidFundAdmin
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
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

func (k msgServer) TransferOwnership(goCtx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	k.SetOwner(ctx, msg.NewOwner)

	return &types.MsgTransferOwnershipResponse{}, ctx.EventManager().EmitTypedEvent(&types.OwnershipTransferred{
		PreviousOwner: owner,
		NewOwner:      msg.NewOwner,
	})
}
