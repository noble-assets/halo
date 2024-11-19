// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v3/keeper"
	"github.com/noble-assets/halo/v3/types"
	"github.com/noble-assets/halo/v3/types/entitlements"
	"github.com/noble-assets/halo/v3/utils"
	"github.com/noble-assets/halo/v3/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestPause(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsMsgServer(k)

	// ACT: Attempt to pause with no owner set.
	_, err := server.Pause(ctx, &entitlements.MsgPause{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to pause with invalid signer.
	_, err = server.Pause(ctx, &entitlements.MsgPause{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.Paused
	k.Paused = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		entitlements.PausedKey, "entitlements_paused", collections.BoolValue,
	)

	// ACT: Attempt to pause with failing Paused collection store.
	_, err = server.Pause(ctx, &entitlements.MsgPause{
		Signer: owner.Address,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.Paused = tmp

	// ACT: Attempt to pause.
	_, err = server.Pause(ctx, &entitlements.MsgPause{
		Signer: owner.Address,
	})
	// ASSERT: The action should've succeeded, and set paused state to true.
	require.NoError(t, err)
	require.True(t, k.GetPaused(ctx))
}

func TestUnpause(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsMsgServer(k)

	// ARRANGE: Set paused state to true.
	err := k.SetPaused(ctx, true)
	require.NoError(t, err)

	// ACT: Attempt to unpause with no owner set.
	_, err = server.Unpause(ctx, &entitlements.MsgUnpause{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to unpause with invalid signer.
	_, err = server.Unpause(ctx, &entitlements.MsgUnpause{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.Paused
	k.Paused = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		entitlements.PausedKey, "entitlements_paused", collections.BoolValue,
	)

	// ACT: Attempt to unpause with failing Paused collection store.
	_, err = server.Unpause(ctx, &entitlements.MsgUnpause{
		Signer: owner.Address,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.Paused = tmp

	// ACT: Attempt to unpause.
	_, err = server.Unpause(ctx, &entitlements.MsgUnpause{
		Signer: owner.Address,
	})
	// ASSERT: The action should've succeeded, and set paused state to false.
	require.NoError(t, err)
	require.False(t, k.GetPaused(ctx))
}

func TestEntitlementsTransferOwnership(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsMsgServer(k)

	// ACT: Attempt to transfer ownership with no owner set.
	_, err := server.TransferOwnership(ctx, &entitlements.MsgTransferOwnership{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to transfer ownership with invalid signer.
	_, err = server.TransferOwnership(ctx, &entitlements.MsgTransferOwnership{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt to transfer ownership to same address.
	_, err = server.TransferOwnership(ctx, &entitlements.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: owner.Address,
	})
	// ASSERT: The action should've failed due to same address.
	require.ErrorContains(t, err, entitlements.ErrSameOwner.Error())

	// ARRANGE: Generate a new owner account.
	newOwner := utils.TestAccount()

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.EntitlementsOwner
	k.EntitlementsOwner = collections.NewItem(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		entitlements.OwnerKey, "entitlements_owner", collections.StringValue,
	)

	// ACT: Attempt to transfer ownership with failing EntitlementsOwner collection store.
	_, err = server.TransferOwnership(ctx, &entitlements.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.EntitlementsOwner = tmp

	// ACT: Attempt to transfer ownership.
	_, err = server.TransferOwnership(ctx, &entitlements.MsgTransferOwnership{
		Signer:   owner.Address,
		NewOwner: newOwner.Address,
	})
	// ASSERT: The action should've succeeded, and set owner in state.
	require.NoError(t, err)
	require.Equal(t, newOwner.Address, k.GetEntitlementsOwner(ctx))
}

func TestEntitlementsPublicCapabilities(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsMsgServer(k)

	// ASSERT: Initial capability roles genesis state
	require.Equal(t, 24, len(k.GetAllCapabilityRoles(ctx)))

	// ACT: Attempt to set public capability with no owner set.
	_, err := server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt set public capability with invalid signer.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt set public capability with invalid method.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer:  owner.Address,
		Method:  "transfer2",
		Enabled: true,
	})
	// ASSERT: The action should've failed due to invalid method.
	require.ErrorContains(t, err, entitlements.ErrInvalidMethod.Error())

	// ACT: Attempt set public capability with non-allowed method.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer:  owner.Address,
		Method:  "/cosmos.bank.v1beta1.MsgSend",
		Enabled: true,
	})
	// ASSERT: The action should've failed due to non-allowed cosmos method.
	require.ErrorContains(t, err, entitlements.ErrInvalidMethod.Error())

	// ACT: Attempt set public capability a valid capability.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer:  owner.Address,
		Method:  "/halo.entitlements.v1.MsgSetRoleCapability",
		Enabled: true,
	})
	// ASSERT: The action should've succeeded, and set method and enabled state.
	require.NoError(t, err)
	require.Equal(t, true, k.IsPublicCapability(ctx, "/halo.entitlements.v1.MsgSetRoleCapability"))
	require.Equal(t, 1, len(k.GetPublicCapabilities(ctx)))

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.PublicCapabilities
	k.PublicCapabilities = collections.NewMap(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		entitlements.PublicPrefix, "entitlements_public_capabilities", collections.StringKey, collections.BoolValue,
	)

	// ACT: Attempt set public capability with failing PublicCapabilities collection store.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Enabled: true,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.PublicCapabilities = tmp

	// ACT: Attempt set public capability a valid capability.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Enabled: true,
	})
	// ASSERT: The action should've succeeded, and set method and enabled state.
	require.NoError(t, err)
	require.Equal(t, true, k.IsPublicCapability(ctx, "transfer"))
	require.Equal(t, 2, len(k.GetPublicCapabilities(ctx)))

	// ACT: Attempt to update a public capability.
	_, err = server.SetPublicCapability(ctx, &entitlements.MsgSetPublicCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Enabled: false,
	})
	// ASSERT: The action should've succeeded, and set method and enabled state.
	require.NoError(t, err)
	require.Equal(t, false, k.IsPublicCapability(ctx, "transfer"))
}

func TestEntitlementsUserRoles(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsMsgServer(k)

	// ACT: Attempt to set user role with no owner set.
	_, err := server.SetUserRole(ctx, &entitlements.MsgSetUserRole{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner, bob, alice := utils.TestAccount(), utils.TestAccount(), utils.TestAccount()
	userAddress, _ := sdk.AccAddressFromBech32(bob.Address)
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt set user role with invalid signer.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt set user role with an invalid user address.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Invalid,
		Role:    2,
		Enabled: true,
	})
	// ASSERT: The action should've failed due to invalid address.
	require.Error(t, err)

	// ACT: Attempt set user role with a negative invalid role.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Address,
		Role:    -2,
		Enabled: true,
	})
	// ASSERT: The action should've failed due to an invalid negative role.
	require.ErrorContains(t, err, entitlements.ErrInvalidRole.Error())

	// ACT: Attempt set user role with a non-existing role.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Address,
		Role:    100,
		Enabled: true,
	})
	// ASSERT: The action should've failed due to a non-existing role.
	require.ErrorContains(t, err, entitlements.ErrInvalidRole.Error())

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmp := k.UserRoles
	k.UserRoles = collections.NewMap(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		entitlements.UserPrefix, "entitlements_user_roles", collections.BytesKey, collections.BoolValue,
	)

	// ACT: Attempt set user role with failing UserRoles collection store.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Address,
		Role:    entitlements.ROLE_INTERNATIONAL_FEEDER,
		Enabled: true,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.UserRoles = tmp

	// ACT: Attempt set user role with valid message.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Address,
		Role:    entitlements.ROLE_INTERNATIONAL_FEEDER,
		Enabled: true,
	})
	// ASSERT: The action should've succeeded, and set the role.
	require.NoError(t, err)
	require.Equal(t, 1, len(k.GetUserRoles(ctx, userAddress)))

	// ACT: Attempt to add additional role to user.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Address,
		Role:    entitlements.ROLE_LIQUIDITY_PROVIDER,
		Enabled: true,
	})
	// ASSERT: The action should've succeeded, and set the role.
	require.NoError(t, err)
	require.Equal(t, 2, len(k.GetUserRoles(ctx, userAddress)))

	// ACT: Attempt to disable the user role.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    bob.Address,
		Role:    entitlements.ROLE_INTERNATIONAL_FEEDER,
		Enabled: false,
	})
	// ASSERT: The action should've succeeded, and set the role.
	require.NoError(t, err)
	require.Equal(t, []entitlements.UserRole{
		{User: bob.Address, Role: 2, Enabled: false},
		{User: bob.Address, Role: 6, Enabled: true},
	}, k.GetAllUserRoles(ctx))

	// ACT: Attempt to disable the user role.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer:  owner.Address,
		User:    alice.Address,
		Role:    entitlements.ROLE_INTERNATIONAL_SDYF,
		Enabled: false,
	})
	// ASSERT: The action should've succeeded, and set the role.
	require.NoError(t, err)
	require.ElementsMatch(t, []entitlements.UserRole{
		{User: bob.Address, Role: 2, Enabled: false},
		{User: bob.Address, Role: 6, Enabled: true},
		{User: alice.Address, Role: 4, Enabled: false},
	}, k.GetAllUserRoles(ctx))
}

func TestEntitlementsUserCapability(t *testing.T) {
	k, ctx := mocks.HaloKeeper(t)
	server := keeper.NewEntitlementsMsgServer(k)

	// ASSERT: Initial genesis capabilities state
	require.Equal(t, 4, len(k.GetCapabilityRoles(ctx, "transfer")))

	// ACT: Attempt to set role capability with no owner set.
	_, err := server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{})
	// ASSERT: The action should've failed due to no owner set.
	require.ErrorContains(t, err, "there is no owner")

	// ARRANGE: Set entitlements owner in state.
	owner := utils.TestAccount()
	err = k.SetEntitlementsOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt set role capability with invalid signer.
	_, err = server.SetUserRole(ctx, &entitlements.MsgSetUserRole{
		Signer: utils.TestAccount().Address,
	})
	// ASSERT: The action should've failed due to invalid signer.
	require.ErrorContains(t, err, entitlements.ErrInvalidOwner.Error())

	// ACT: Attempt set role capability with a negative invalid role.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Role:    -1000,
		Method:  "transfer",
		Enabled: false,
	})
	// ASSERT: The action should've failed due to an invalid negative role.
	require.ErrorContains(t, err, entitlements.ErrInvalidRole.Error())

	// ACT: Attempt set role capability with a non-existing role.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Role:    100,
		Enabled: true,
	})
	// ASSERT: The action should've failed due to a non-existing role.
	require.ErrorContains(t, err, entitlements.ErrInvalidRole.Error())

	// ACT: Attempt set role capability with a non-existing role.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Role:    100,
		Enabled: false,
	})
	// ASSERT: The action should've failed due to a non-existing role.
	require.ErrorContains(t, err, entitlements.ErrInvalidRole.Error())

	require.Equal(t, 4, len(k.GetCapabilityRoles(ctx, "transfer")))
	// ACT: Attempt set role capability with a non-existing method.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "transfer2",
		Role:    2,
		Enabled: false,
	})
	// ASSERT: The action should've failed due to a non-existing role.
	require.ErrorContains(t, err, entitlements.ErrInvalidMethod.Error())

	// ACT: Attempt set role capability with a non-allowed method.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "/cosmos.bank.v1beta1.MsgSend",
		Role:    2,
		Enabled: false,
	})
	// ASSERT: The action should've failed due to a non-allowed role.
	require.ErrorContains(t, err, entitlements.ErrInvalidMethod.Error())

	// ARRANGE: Set up a failing collection store for the attribute setter.
	tmpRole := k.RoleCapabilities
	k.RoleCapabilities = collections.NewMap(
		collections.NewSchemaBuilder(mocks.FailingStore(mocks.Set, utils.GetKVStore(ctx, types.ModuleName))),
		entitlements.CapabilityPrefix, "entitlements_role_capabilities", collections.BytesKey, collections.BoolValue,
	)

	// ACT: Attempt set role capability with failing RoleCapabilities collection store.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Role:    5,
		Enabled: true,
	})
	// ASSERT: The action should've failed due to collection store setter error.
	require.Error(t, err, mocks.ErrorStoreAccess)
	k.RoleCapabilities = tmpRole

	// ACT: Attempt set role capability with valid message.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Role:    5,
		Enabled: true,
	})
	// ASSERT: The action should've succeeded, and set the role capability.
	require.NoError(t, err)
	require.Equal(t, 5, len(k.GetCapabilityRoles(ctx, "transfer")))

	// ACT: Attempt remove role capability.
	_, err = server.SetRoleCapability(ctx, &entitlements.MsgSetRoleCapability{
		Signer:  owner.Address,
		Method:  "transfer",
		Role:    5,
		Enabled: false,
	})
	// ASSERT: The action should've succeeded, and removed the role capability.
	require.NoError(t, err)
	require.Equal(t, 4, len(k.GetCapabilityRoles(ctx, "transfer")))
}
