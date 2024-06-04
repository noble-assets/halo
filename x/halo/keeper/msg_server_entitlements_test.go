package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/stretchr/testify/require"
)

func TestPause(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsMsgServer(k)

	// ACT: Attempt to pause with no owner set.
	_, err := server.Pause(goCtx, &entitlements.MsgPause{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	k.SetEntitlementsOwner(ctx, owner.Address)

	// ACT: Attempt to pause with invalid signer.
	_, err = server.Pause(goCtx, &entitlements.MsgPause{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt to pause.
	_, err = server.Pause(goCtx, &entitlements.MsgPause{
		Signer: owner.Address,
	})
	// ASSERT: The action should've succeeded, and set paused state to true.
	require.NoError(t, err)
	require.True(t, k.GetPaused(ctx))
}

func TestUnpause(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsMsgServer(k)

	// ARRANGE: Set paused state to true.
	k.SetPaused(ctx, true)

	// ACT: Attempt to unpause with no owner set.
	_, err := server.Unpause(goCtx, &entitlements.MsgUnpause{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	k.SetEntitlementsOwner(ctx, owner.Address)

	// ACT: Attempt to unpause with invalid signer.
	_, err = server.Unpause(goCtx, &entitlements.MsgUnpause{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt to unpause.
	_, err = server.Unpause(goCtx, &entitlements.MsgUnpause{
		Signer: owner.Address,
	})
	// ASSERT: The action should've succeeded, and set paused state to false.
	require.NoError(t, err)
	require.False(t, k.GetPaused(ctx))
}

func TestEntitlementsTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(goCtx, &entitlements.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	k.SetEntitlementsOwner(ctx, owner.Address)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(goCtx, &entitlements.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(goCtx, &entitlements.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, entitlements.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(goCtx, &entitlements.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetEntitlementsOwner(ctx))
}
