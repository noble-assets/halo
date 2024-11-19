// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noble-assets/halo/v3/keeper"
	"github.com/noble-assets/halo/v3/types/entitlements"
	"github.com/noble-assets/halo/v3/utils"
	"github.com/noble-assets/halo/v3/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestEntitlementsOwnerQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query entitlements owner with invalid request.
	_, err := server.Owner(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to query entitlements owner.
	res, err := server.Owner(ctx, &entitlements.QueryOwner{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
}

func TestPausedQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query paused state with invalid request.
	_, err := server.Paused(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query paused state with no state.
	res, err := server.Paused(ctx, &entitlements.QueryPaused{})
	// ASSERT: The query should've succeeded, and returned false.
	require.NoError(t, err)
	require.False(t, res.Paused)

	// ARRANGE: Set paused state to true.
	err = k.SetPaused(ctx, true)
	require.NoError(t, err)

	// ACT: Attempt to query paused state with state.
	res, err = server.Paused(ctx, &entitlements.QueryPaused{})
	// ASSERT: The query should've succeeded, and returned true.
	require.NoError(t, err)
	require.True(t, res.Paused)
}

func TestPublicCapabilityQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query public capability with invalid request.
	_, err := server.PublicCapability(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query public capability with invalid method.
	_, err = server.PublicCapability(ctx, &entitlements.QueryPublicCapability{
		Method: "",
	})
	// ASSERT: The query should've failed due to invalid method.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query public capability that is disabled.
	res, err := server.PublicCapability(ctx, &entitlements.QueryPublicCapability{
		Method: "transfer",
	})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.False(t, res.Enabled)

	// ARRANGE: Set public capability in state.
	// NOTE: Transferring will never be public, this is just for testing.
	err = k.SetPublicCapability(ctx, "transfer", true)
	require.NoError(t, err)

	// ACT: Attempt to query public capability that is enabled.
	res, err = server.PublicCapability(ctx, &entitlements.QueryPublicCapability{
		Method: "transfer",
	})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.True(t, res.Enabled)
}

func TestRoleCapabilityQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query role capability with invalid request.
	_, err := server.RoleCapability(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query role capability with invalid method.
	_, err = server.RoleCapability(ctx, &entitlements.QueryRoleCapability{
		Method: "",
	})
	// ASSERT: The query should've failed due to invalid method.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query role capability with an invalid method.
	_, err = server.RoleCapability(ctx, &entitlements.QueryRoleCapability{
		Method: "",
	})
	// ASSERT: The query should've failed to invalid method.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query role capability with a non-existing method.
	res, err := server.RoleCapability(ctx, &entitlements.QueryRoleCapability{
		Method: "non-existing",
	})
	// ASSERT: The query should've succeeded without results.
	require.NoError(t, err)
	require.Equal(t, 0, len(res.Roles))

	// ACT: Attempt to query role capability with an existing method.
	res, err = server.RoleCapability(ctx, &entitlements.QueryRoleCapability{
		Method: "transfer",
	})
	// ASSERT: The query should've succeeded with results.
	require.NoError(t, err)
	require.Equal(t, 4, len(res.Roles))
}

func TestUserCapabilityQuery(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsQueryServer(k)

	// ACT: Attempt to query user capability with invalid request.
	_, err := server.UserCapability(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query user capability with an empty address.
	_, err = server.UserCapability(ctx, &entitlements.QueryUserCapability{
		Address: "",
	})
	// ASSERT: The query should've failed due to empty address.
	require.ErrorContains(t, err, "unable to decode address")

	user := utils.TestAccount()

	// ACT: Attempt to user capability with an invalid address.
	_, err = server.UserCapability(ctx, &entitlements.QueryUserCapability{
		Address: user.Invalid,
	})
	// ASSERT: The query should've failed to invalid address.
	require.ErrorContains(t, err, "unable to decode address")

	// ACT: Attempt to query user capability with an address without capabilities.
	res, err := server.UserCapability(ctx, &entitlements.QueryUserCapability{
		Address: user.Address,
	})
	// ASSERT: The query should've succeeded without results.
	require.NoError(t, err)
	require.Equal(t, 0, len(res.Roles))

	userAddress, _ := sdk.AccAddressFromBech32(user.Address)
	err = k.SetUserRole(ctx, userAddress, 2, true)
	require.NoError(t, err)

	// ACT: Attempt to query role capability with an existing user and capability.
	res, err = server.UserCapability(ctx, &entitlements.QueryUserCapability{
		Address: user.Address,
	})
	// ASSERT: The query should've succeeded with results.
	require.NoError(t, err)
	require.Equal(t, 1, len(res.Roles))
	require.ElementsMatch(t, []entitlements.UserRole{
		{User: user.Address, Role: 2, Enabled: true},
	}, k.GetAllUserRoles(ctx))

	err = k.SetUserRole(ctx, userAddress, 3, true)
	require.NoError(t, err)
	// ACT: Attempt to query role capability with an existing user and multiple capabilities.
	res, err = server.UserCapability(ctx, &entitlements.QueryUserCapability{
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
