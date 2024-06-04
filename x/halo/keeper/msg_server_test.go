package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/stretchr/testify/require"
)

var ONE = sdk.NewInt(1_000_000)

func TestBurn(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate a user account.
	user := utils.TestAccount()

	// ACT: Attempt to burn with an invalid signer address.
	_, err := server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to burn without required permissions.
	_, err = server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot execute /halo.v1.MsgBurn")

	// ARRANGE: Assign the international feeder role to user.
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to burn with insufficient funds.
	_, err = server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 1 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	// ACT: Attempt to burn.
	_, err = server.Burn(goCtx, &types.MsgBurn{
		Signer: user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestBurnFor(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to burn for with an invalid signer address.
	_, err := server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to burn for with an invalid signer.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to burn for with an invalid from address.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid from address.
	require.ErrorContains(t, err, "unable to decode from address")

	// ACT: Attempt to burn for with insufficient funds.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "unable to transfer from account to module")

	// ARRANGE: Give user 1 $USYC.
	bank.Balances[user.Address] = sdk.NewCoins(sdk.NewCoin(k.Denom, ONE))

	// ACT: Attempt to burn.
	_, err = server.BurnFor(goCtx, &types.MsgBurnFor{
		Signer: admin.Address,
		From:   user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.True(t, bank.Balances[user.Address].IsZero())
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestMint(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate admin and user accounts.
	admin, user := utils.TestAccount(), utils.TestAccount()
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)
	k.SetUserRole(ctx, user.Bytes, entitlements.ROLE_INTERNATIONAL_FEEDER, true)

	// ACT: Attempt to mint with an invalid signer address.
	_, err := server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to mint with an invalid signer.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ACT: Attempt to mint with an invalid to address.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Invalid,
	})
	// ASSERT: The action should've failed due to invalid to address.
	require.ErrorContains(t, err, "unable to decode to address")

	// ACT: Attempt to mint without required permissions.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Address,
		To:     utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid permissions.
	require.ErrorContains(t, err, "cannot transfer")

	// ACT: Attempt to mint.
	_, err = server.Mint(goCtx, &types.MsgMint{
		Signer: admin.Address,
		To:     user.Address,
		Amount: ONE,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, ONE, bank.Balances[user.Address].AmountOf(k.Denom))
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestTradeToFiat(t *testing.T) {
	bank := mocks.BankKeeper{
		Balances:    make(map[string]sdk.Coins),
		Restriction: mocks.NoOpSendRestrictionFn,
	}
	k, ctx := mocks.HaloKeeperWithKeepers(t, mocks.AccountKeeper{}, bank)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ARRANGE: Generate an admin account.
	admin := utils.TestAccount()

	// ACT: Attempt to trade to fiat with an invalid signer address.
	_, err := server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid signer address.
	require.ErrorContains(t, err, "unable to decode signer address")

	// ACT: Attempt to trade to fiat with an invalid signer.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer: admin.Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidFundAdmin.Error())

	// ARRANGE: Set fund admin in state.
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_FUND_ADMIN, true)

	// ACT: Attempt to trade to fiat with an invalid recipient address.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Recipient: admin.Invalid,
	})
	// ASSERT: The action should've failed due to invalid recipient address.
	require.ErrorContains(t, err, "unable to decode recipient address")

	// ACT: Attempt to trade to fiat with invalid recipient permissions.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Recipient: admin.Address,
	})
	// ASSERT: The action should've failed due to invalid recipient permissions.
	require.ErrorContains(t, err, types.ErrInvalidLiquidityProvider.Error())

	// ARRANGE: Set liquidity provider in state.
	k.SetUserRole(ctx, admin.Bytes, entitlements.ROLE_LIQUIDITY_PROVIDER, true)

	// ACT: Attempt to trade to fiat with insufficient funds.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    ONE,
		Recipient: admin.Address,
	})
	// ASSERT: The action should've failed due to insufficient funds.
	require.ErrorContains(t, err, "insufficient funds")

	// ARRANGE: Give the module 1 $USDC.
	bank.Balances[types.ModuleAddress.String()] = sdk.NewCoins(sdk.NewCoin(k.Underlying, ONE))

	// ACT: Attempt to trade to fiat.
	_, err = server.TradeToFiat(goCtx, &types.MsgTradeToFiat{
		Signer:    admin.Address,
		Amount:    ONE,
		Recipient: admin.Address,
	})
	// ASSERT: The action should've succeeded.
	require.NoError(t, err)
	require.Equal(t, ONE, bank.Balances[admin.Address].AmountOf(k.Underlying))
	require.True(t, bank.Balances[types.ModuleName].IsZero())
}

func TestTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(goCtx, &types.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, types.ErrNoOwner.Error())

	// ARRANGE: Set owner in state.
	owner := utils.TestAccount()
	k.SetOwner(ctx, owner.Address)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(goCtx, &types.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, types.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(goCtx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, types.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(goCtx, &types.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetOwner(ctx))
}
