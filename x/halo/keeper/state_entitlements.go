package keeper

import (
	"encoding/binary"
	"slices"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

func (k *Keeper) CanCall(ctx sdk.Context, user []byte, method string) bool {
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

func (k *Keeper) GetEntitlementsOwner(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(entitlements.OwnerKey))
}

func (k *Keeper) SetEntitlementsOwner(ctx sdk.Context, owner string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(entitlements.OwnerKey, []byte(owner))
}

//

func (k *Keeper) GetPaused(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(entitlements.PausedKey)
	if len(bz) == 1 && bz[0] == 1 {
		return true
	} else {
		return false
	}
}

func (k *Keeper) SetPaused(ctx sdk.Context, paused bool) {
	store := ctx.KVStore(k.storeKey)
	if paused {
		store.Set(entitlements.PausedKey, []byte{0x1})
	} else {
		store.Set(entitlements.PausedKey, []byte{0x0})
	}
}

//

func (k *Keeper) IsPublicCapability(ctx sdk.Context, method string) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(entitlements.PublicKey(method))
	if len(bz) == 1 && bz[0] == 1 {
		return true
	} else {
		return false
	}
}

func (k *Keeper) GetPublicCapabilities(ctx sdk.Context) (publicCapabilities map[string]bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), entitlements.PublicPrefix)
	itr := store.Iterator(nil, nil)

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		enabled := false
		if len(itr.Value()) == 1 && itr.Value()[0] == 1 {
			enabled = true
		}

		publicCapabilities[string(itr.Key())] = enabled
	}

	return
}

func (k *Keeper) SetPublicCapability(ctx sdk.Context, method string, enabled bool) {
	store := ctx.KVStore(k.storeKey)
	key := entitlements.PublicKey(method)
	if enabled {
		store.Set(key, []byte{0x1})
	} else {
		store.Set(key, []byte{0x0})
	}
}

//

func (k *Keeper) GetCapabilityRoles(ctx sdk.Context, method string) (roles []entitlements.Role) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), entitlements.CapabilityKey(method))
	itr := store.Iterator(nil, nil)

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		role := entitlements.Role(binary.BigEndian.Uint64(itr.Key()))
		enabled := len(itr.Value()) == 1 && itr.Value()[0] == 1

		if role != entitlements.ROLE_UNSPECIFIED && enabled {
			roles = append(roles, role)
		}
	}

	return
}

func (k *Keeper) GetAllCapabilityRoles(ctx sdk.Context) (capabilityRoles []entitlements.RoleCapability) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), entitlements.CapabilityPrefix)
	itr := store.Iterator(nil, nil)

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		enabled := false
		if len(itr.Value()) == 1 && itr.Value()[0] == 1 {
			enabled = true
		}

		capabilityRoles = append(capabilityRoles, entitlements.RoleCapability{
			Method:  string(itr.Key()[:len(itr.Key())-8]),
			Role:    entitlements.Role(binary.BigEndian.Uint64(itr.Key()[len(itr.Key())-8:])),
			Enabled: enabled,
		})
	}

	return
}

func (k *Keeper) SetRoleCapability(ctx sdk.Context, method string, role entitlements.Role, enabled bool) {
	store := ctx.KVStore(k.storeKey)
	key := entitlements.CapabilityRoleKey(method, role)
	if enabled {
		store.Set(key, []byte{0x1})
	} else {
		store.Set(key, []byte{0x0})
	}
}

//

func (k *Keeper) GetUserRoles(ctx sdk.Context, user []byte) (roles []entitlements.Role) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), entitlements.UserKey(user))
	itr := store.Iterator(nil, nil)

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		role := entitlements.Role(binary.BigEndian.Uint64(itr.Key()))
		enabled := len(itr.Value()) == 1 && itr.Value()[0] == 1

		if role != entitlements.ROLE_UNSPECIFIED && enabled {
			roles = append(roles, role)
		}
	}

	return
}

func (k *Keeper) GetAllUserRoles(ctx sdk.Context) (userRoles []entitlements.UserRole) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), entitlements.UserPrefix)
	itr := store.Iterator(nil, nil)

	defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		enabled := false
		if len(itr.Value()) == 1 && itr.Value()[0] == 1 {
			enabled = true
		}

		userRoles = append(userRoles, entitlements.UserRole{
			User:    sdk.AccAddress(itr.Key()[:len(itr.Key())-8]).String(),
			Role:    entitlements.Role(binary.BigEndian.Uint64(itr.Key()[len(itr.Key())-8:])),
			Enabled: enabled,
		})
	}

	return
}

func (k *Keeper) HasRole(ctx sdk.Context, address []byte, role entitlements.Role) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(entitlements.UserRoleKey(address, role))
	if len(bz) == 1 && bz[0] == 1 {
		return true
	} else {
		return false
	}
}

func (k *Keeper) SetUserRole(ctx sdk.Context, user []byte, role entitlements.Role, enabled bool) {
	store := ctx.KVStore(k.storeKey)
	key := entitlements.UserRoleKey(user, role)
	if enabled {
		store.Set(key, []byte{0x1})
	} else {
		store.Set(key, []byte{0x0})
	}
}
