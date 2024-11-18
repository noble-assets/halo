// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"
	"encoding/binary"
	"slices"

	"cosmossdk.io/collections"
	"github.com/noble-assets/halo/v3/types/entitlements"
)

func (k *Keeper) CanCall(ctx context.Context, user []byte, method string) bool {
	if k.GetPaused(ctx) {
		return false
	}

	userRoles := k.GetUserRoles(ctx, user)
	capabilityRoles := k.GetCapabilityRoles(ctx, method)

	for _, role := range capabilityRoles {
		if slices.Contains(userRoles, role) {
			return true
		}
	}

	return k.IsPublicCapability(ctx, method)
}

//

func (k *Keeper) GetEntitlementsOwner(ctx context.Context) string {
	owner, _ := k.EntitlementsOwner.Get(ctx)
	return owner
}

func (k *Keeper) SetEntitlementsOwner(ctx context.Context, owner string) error {
	return k.EntitlementsOwner.Set(ctx, owner)
}

//

func (k *Keeper) GetPaused(ctx context.Context) bool {
	paused, _ := k.Paused.Get(ctx)
	return paused
}

func (k *Keeper) SetPaused(ctx context.Context, paused bool) error {
	return k.Paused.Set(ctx, paused)
}

//

func (k *Keeper) IsPublicCapability(ctx context.Context, method string) bool {
	enabled, _ := k.PublicCapabilities.Get(ctx, method)
	return enabled
}

func (k *Keeper) GetPublicCapabilities(ctx context.Context) map[string]bool {
	publicCapabilities := make(map[string]bool)

	_ = k.PublicCapabilities.Walk(ctx, nil, func(method string, enabled bool) (stop bool, err error) {
		publicCapabilities[method] = enabled
		return false, nil
	})

	return publicCapabilities
}

func (k *Keeper) SetPublicCapability(ctx context.Context, method string, enabled bool) error {
	return k.PublicCapabilities.Set(ctx, method, enabled)
}

//

func (k *Keeper) GetCapabilityRoles(ctx context.Context, method string) []entitlements.Role {
	var roles []entitlements.Role

	itr, _ := k.RoleCapabilities.Iterate(ctx, new(collections.Range[[]byte]).Prefix([]byte(method)))

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		key, _ := itr.Key()
		enabled, _ := itr.Value()

		if enabled {
			role := binary.BigEndian.Uint64(key[len(key)-8:])
			roles = append(roles, entitlements.Role(role))
		}
	}

	return roles
}

func (k *Keeper) GetAllCapabilityRoles(ctx context.Context) []entitlements.RoleCapability {
	var capabilityRoles []entitlements.RoleCapability

	_ = k.RoleCapabilities.Walk(ctx, nil, func(key []byte, enabled bool) (stop bool, err error) {
		capabilityRoles = append(capabilityRoles, entitlements.RoleCapability{
			Method:  string(key[:len(key)-8]),
			Role:    entitlements.Role(binary.BigEndian.Uint64(key[len(key)-8:])),
			Enabled: enabled,
		})

		return false, nil
	})

	return capabilityRoles
}

func (k *Keeper) SetRoleCapability(ctx context.Context, method string, role entitlements.Role, enabled bool) error {
	return k.RoleCapabilities.Set(ctx, entitlements.CapabilityRoleKey(method, role), enabled)
}

//

func (k *Keeper) GetUserRoles(ctx context.Context, address []byte) []entitlements.Role {
	var roles []entitlements.Role

	itr, _ := k.UserRoles.Iterate(ctx, new(collections.Range[[]byte]).Prefix(address))

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		key, _ := itr.Key()
		enabled, _ := itr.Value()

		if enabled {
			role := binary.BigEndian.Uint64(key[len(key)-8:])
			roles = append(roles, entitlements.Role(role))
		}
	}

	return roles
}

func (k *Keeper) GetAllUserRoles(ctx context.Context) []entitlements.UserRole {
	var userRoles []entitlements.UserRole

	_ = k.UserRoles.Walk(ctx, nil, func(key []byte, enabled bool) (stop bool, err error) {
		address, _ := k.addressCodec.BytesToString(key[:len(key)-8])
		userRoles = append(userRoles, entitlements.UserRole{
			User:    address,
			Role:    entitlements.Role(binary.BigEndian.Uint64(key[len(key)-8:])),
			Enabled: enabled,
		})

		return false, nil
	})

	return userRoles
}

func (k *Keeper) HasRole(ctx context.Context, address []byte, role entitlements.Role) bool {
	enabled, _ := k.UserRoles.Get(ctx, entitlements.UserRoleKey(address, role))
	return enabled
}

func (k *Keeper) SetUserRole(ctx context.Context, address []byte, role entitlements.Role, enabled bool) error {
	return k.UserRoles.Set(ctx, entitlements.UserRoleKey(address, role), enabled)
}
