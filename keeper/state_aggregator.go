// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/noble-assets/halo/v3/types/aggregator"
)

//

func (k *Keeper) GetAggregatorOwner(ctx context.Context) string {
	owner, _ := k.AggregatorOwner.Get(ctx)
	return owner
}

func (k *Keeper) SetAggregatorOwner(ctx context.Context, owner string) error {
	return k.AggregatorOwner.Set(ctx, owner)
}

//

func (k *Keeper) GetLastRoundId(ctx context.Context) uint64 {
	id, _ := k.LastRoundId.Peek(ctx)
	return id
}

func (k *Keeper) IncrementLastRoundId(ctx context.Context) (uint64, error) {
	return k.LastRoundId.Next(ctx)
}

func (k *Keeper) SetLastRoundId(ctx context.Context, id uint64) error {
	return k.LastRoundId.Set(ctx, id)
}

//

func (k *Keeper) GetNextPrice(ctx context.Context) math.Int {
	price, _ := k.NextPrice.Get(ctx)
	return price
}

func (k *Keeper) SetNextPrice(ctx context.Context, price math.Int) error {
	return k.NextPrice.Set(ctx, price)
}

//

func (k *Keeper) GetRound(ctx context.Context, id uint64) (aggregator.RoundData, bool) {
	round, err := k.Rounds.Get(ctx, id)
	if err != nil {
		return aggregator.RoundData{}, false
	}

	return round, true
}

func (k *Keeper) GetRounds(ctx context.Context) map[uint64]aggregator.RoundData {
	rounds := make(map[uint64]aggregator.RoundData)

	_ = k.Rounds.Walk(ctx, nil, func(id uint64, round aggregator.RoundData) (stop bool, err error) {
		rounds[id] = round
		return false, nil
	})

	return rounds
}

func (k *Keeper) SetRound(ctx context.Context, id uint64, round aggregator.RoundData) error {
	return k.Rounds.Set(ctx, id, round)
}
