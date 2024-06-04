package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types"
)

//

func (k *Keeper) GetOwner(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.OwnerKey))
}

func (k *Keeper) SetOwner(ctx sdk.Context, owner string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.OwnerKey, []byte(owner))
}

//

func (k *Keeper) GetNonce(ctx sdk.Context, address []byte) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NonceKey(address))
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k *Keeper) GetNonces(ctx sdk.Context) map[string]uint64 {
	store := ctx.KVStore(k.storeKey)
	itr := sdk.KVStorePrefixIterator(store, types.NoncePrefix)

	defer itr.Close()

	nonces := make(map[string]uint64)
	for ; itr.Valid(); itr.Next() {
		nonces[sdk.AccAddress(itr.Key()).String()] = binary.BigEndian.Uint64(itr.Value())
	}

	return nonces
}

func (k *Keeper) IncrementNonce(ctx sdk.Context, address []byte) uint64 {
	nonce := k.GetNonce(ctx, address)
	k.SetNonce(ctx, address, nonce+1)
	return nonce
}

func (k *Keeper) SetNonce(ctx sdk.Context, address []byte, nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, nonce)
	store.Set(types.NonceKey(address), bz)
}
