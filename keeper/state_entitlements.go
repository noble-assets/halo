package keeper

import (
	"context"
	"slices"

	"cosmossdk.io/collections"
	"github.com/noble-assets/halo/v2/types/entitlements"
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

func (k *Keeper) SetEntitlementsOwner(ctx context.Context, owner string) {
	_ = k.EntitlementsOwner.Set(ctx, owner)
}

//

func (k *Keeper) GetPaused(ctx context.Context) bool {
	paused, _ := k.Paused.Get(ctx)
	return paused
}

func (k *Keeper) SetPaused(ctx context.Context, paused bool) {
	_ = k.Paused.Set(ctx, paused)
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

func (k *Keeper) SetPublicCapability(ctx context.Context, method string, enabled bool) {
	_ = k.PublicCapabilities.Set(ctx, method, enabled)
}

//

func (k *Keeper) GetCapabilityRoles(ctx context.Context, method string) []entitlements.Role {
	var roles []entitlements.Role

	_ = k.RoleCapabilities.Walk(ctx, collections.NewPrefixedPairRange[string, uint64](method), func(key collections.Pair[string, uint64], enabled bool) (stop bool, err error) {
		if enabled {
			roles = append(roles, entitlements.Role(key.K2()))
		}

		return false, nil
	})

	return roles
}

func (k *Keeper) GetAllCapabilityRoles(ctx context.Context) []entitlements.RoleCapability {
	var capabilityRoles []entitlements.RoleCapability

	_ = k.RoleCapabilities.Walk(ctx, nil, func(key collections.Pair[string, uint64], enabled bool) (stop bool, err error) {
		capabilityRoles = append(capabilityRoles, entitlements.RoleCapability{
			Method:  key.K1(),
			Role:    entitlements.Role(key.K2()),
			Enabled: enabled,
		})

		return false, nil
	})

	return capabilityRoles
}

func (k *Keeper) SetRoleCapability(ctx context.Context, method string, role entitlements.Role, enabled bool) {
	_ = k.RoleCapabilities.Set(ctx, collections.Join(method, uint64(role)), enabled)
}

//

func (k *Keeper) GetUserRoles(ctx context.Context, user []byte) []entitlements.Role {
	var roles []entitlements.Role

	_ = k.UserRoles.Walk(ctx, collections.NewPrefixedPairRange[[]byte, uint64](user), func(key collections.Pair[[]byte, uint64], enabled bool) (stop bool, err error) {
		if enabled {
			roles = append(roles, entitlements.Role(key.K2()))
		}

		return false, nil
	})

	return roles
}

func (k *Keeper) GetAllUserRoles(ctx context.Context) []entitlements.UserRole {
	var userRoles []entitlements.UserRole

	_ = k.UserRoles.Walk(ctx, nil, func(key collections.Pair[[]byte, uint64], enabled bool) (stop bool, err error) {
		address, _ := k.accountKeeper.AddressCodec().BytesToString(key.K1())
		userRoles = append(userRoles, entitlements.UserRole{
			User:    address,
			Role:    entitlements.Role(key.K2()),
			Enabled: enabled,
		})

		return false, nil
	})

	return userRoles
}

func (k *Keeper) HasRole(ctx context.Context, address []byte, role entitlements.Role) bool {
	enabled, _ := k.UserRoles.Get(ctx, collections.Join(address, uint64(role)))
	return enabled
}

func (k *Keeper) SetUserRole(ctx context.Context, address []byte, role entitlements.Role, enabled bool) {
	_ = k.UserRoles.Set(ctx, collections.Join(address, uint64(role)), enabled)
}
