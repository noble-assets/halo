package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/utils"
	"github.com/noble-assets/halo/utils/mocks"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/stretchr/testify/require"
)

func TestEntitlementsOwnerQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query entitlements owner with invalid request.
	_, err := server.Owner(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	k.SetEntitlementsOwner(ctx, owner.Address)

	// ACT: Attempt to query entitlements owner.
	res, err := server.Owner(goCtx, &entitlements.QueryOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
}

func TestPausedQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query paused state with invalid request.
	_, err := server.Paused(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query paused state with no state.
	res, err := server.Paused(goCtx, &entitlements.QueryPaused{})
	// ASSERT: The query should've succeeded, and returned false.
	require.NoError(t, err)
	require.False(t, res.Paused)

	// ARRANGE: Set paused state to true.
	k.SetPaused(ctx, true)

	// ACT: Attempt to query paused state with state.
	res, err = server.Paused(goCtx, &entitlements.QueryPaused{})
	// ASSERT: The query should've succeeded, and returned true.
	require.NoError(t, err)
	require.True(t, res.Paused)
}

func TestPublicCapabilityQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query public capability with invalid request.
	_, err := server.PublicCapability(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query public capability with invalid method.
	_, err = server.PublicCapability(goCtx, &entitlements.QueryPublicCapability{
		Method: "",
	})
	// ASSERT: The query should've failed due to invalid method.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query public capability that is disabled.
	res, err := server.PublicCapability(goCtx, &entitlements.QueryPublicCapability{
		Method: "transfer",
	})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.False(t, res.Enabled)

	// ARRANGE: Set public capability in state.
	// NOTE: Transferring will never be public, this is just for testing.
	k.SetPublicCapability(ctx, "transfer", true)

	// ACT: Attempt to query public capability that is enabled.
	res, err = server.PublicCapability(goCtx, &entitlements.QueryPublicCapability{
		Method: "transfer",
	})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.True(t, res.Enabled)
}

func TestRoleCapabilityQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query role capability with invalid request.
	_, err := server.RoleCapability(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query role capability with invalid method.
	_, err = server.RoleCapability(goCtx, &entitlements.QueryRoleCapability{
		Method: "",
	})
	// ASSERT: The query should've failed due to invalid method.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query role capability with an invalid method.
	res, err := server.RoleCapability(goCtx, &entitlements.QueryRoleCapability{
		Method: "",
	})
	// ASSERT: The query should've failed to invalid method.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query role capability with a non-existing method.
	res, err = server.RoleCapability(goCtx, &entitlements.QueryRoleCapability{
		Method: "non-existing",
	})
	// ASSERT: The query should've succeeded without results.
	require.NoError(t, err)
	require.Equal(t, 0, len(res.Roles))

	// ACT: Attempt to query role capability with an existing method.
	res, err = server.RoleCapability(goCtx, &entitlements.QueryRoleCapability{
		Method: "transfer",
	})
	// ASSERT: The query should've succeeded with results.
	require.NoError(t, err)
	require.Equal(t, 4, len(res.Roles))
}

func TestUserCapabilityQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query user capability with invalid request.
	_, err := server.UserCapability(goCtx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query user capability with an empty address.
	_, err = server.UserCapability(goCtx, &entitlements.QueryUserCapability{
		Address: "",
	})
	// ASSERT: The query should've failed due to empty address.
	require.ErrorContains(t, err, "unable to decode address")

	user := utils.TestAccount()

	// ACT: Attempt to user capability with an invalid address.
	res, err := server.UserCapability(goCtx, &entitlements.QueryUserCapability{
		Address: user.Invalid,
	})
	// ASSERT: The query should've failed to invalid address.
	require.ErrorContains(t, err, "unable to decode address")

	// ACT: Attempt to query user capability with an address without capabilities.
	res, err = server.UserCapability(goCtx, &entitlements.QueryUserCapability{
		Address: user.Address,
	})
	// ASSERT: The query should've succeeded without results.
	require.NoError(t, err)
	require.Equal(t, 0, len(res.Roles))

	userAddress, _ := sdk.AccAddressFromBech32(user.Address)
	k.SetUserRole(ctx, userAddress, 2, true)

	// ACT: Attempt to query role capability with an existing user and capability.
	res, err = server.UserCapability(goCtx, &entitlements.QueryUserCapability{
		Address: user.Address,
	})
	// ASSERT: The query should've succeeded with results.
	require.NoError(t, err)
	require.Equal(t, 1, len(res.Roles))
	require.ElementsMatch(t, []entitlements.UserRole{
		{User: user.Address, Role: 2, Enabled: true},
	}, k.GetAllUserRoles(ctx))

	k.SetUserRole(ctx, userAddress, 3, true)
	// ACT: Attempt to query role capability with an existing user and multiple capabilities.
	res, err = server.UserCapability(goCtx, &entitlements.QueryUserCapability{
		Address: user.Address,
	})
	// ASSERT: The query should've succeeded with results.
	require.NoError(t, err)
	require.Equal(t, 2, len(res.Roles))
	require.ElementsMatch(t, []entitlements.UserRole{
		{User: user.Address, Role: 2, Enabled: true},
		{User: user.Address, Role: 3, Enabled: true},
	}, k.GetAllUserRoles(ctx))
}
