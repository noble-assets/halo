package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
)

//

func (k *Keeper) GetAggregatorOwner(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(aggregator.OwnerKey))
}

func (k *Keeper) SetAggregatorOwner(ctx sdk.Context, owner string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(aggregator.OwnerKey, []byte(owner))
}

//

func (k *Keeper) GetLastRoundId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(aggregator.LastRoundIDKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k *Keeper) IncrementLastRoundId(ctx sdk.Context) uint64 {
	id := k.GetLastRoundId(ctx)
	k.SetLastRoundId(ctx, id+1)
	return id
}

func (k *Keeper) SetLastRoundId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(aggregator.LastRoundIDKey, bz)
}

//

func (k *Keeper) GetNextPrice(ctx sdk.Context) (price sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(aggregator.NextPriceKey)

	_ = price.Unmarshal(bz)
	return
}

func (k *Keeper) SetNextPrice(ctx sdk.Context, price sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz, _ := price.Marshal()
	store.Set(aggregator.NextPriceKey, bz)
}

//

func (k *Keeper) GetRound(ctx sdk.Context, id uint64) (round aggregator.RoundData, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(aggregator.RoundKey(id))
	if bz == nil {
		return round, false
	}

	k.cdc.MustUnmarshal(bz, &round)
	return round, true
}

func (k *Keeper) GetRounds(ctx sdk.Context) map[uint64]aggregator.RoundData {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), aggregator.RoundPrefix)
	itr := store.Iterator(nil, nil)

	defer itr.Close()
	rounds := make(map[uint64]aggregator.RoundData)

	for ; itr.Valid(); itr.Next() {
		id := binary.BigEndian.Uint64(itr.Key())
		var round aggregator.RoundData
		k.cdc.MustUnmarshal(itr.Value(), &round)

		rounds[id] = round
	}

	return rounds
}

func (k *Keeper) SetRound(ctx sdk.Context, id uint64, round aggregator.RoundData) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&round)
	store.Set(aggregator.RoundKey(id), bz)
}
