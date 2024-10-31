// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package keeper

import (
	"context"

	"github.com/noble-assets/halo/v2/types/aggregator"
)

//

func (k *Keeper) GetLastRoundId(ctx context.Context) uint64 {
	id, _ := k.LastRoundId.Peek(ctx)
	return id
}

func (k *Keeper) SetLastRoundId(ctx context.Context, id uint64) error {
	return k.LastRoundId.Set(ctx, id)
}

//

func (k *Keeper) GetRounds(ctx context.Context) map[uint64]aggregator.RoundData {
	rounds := make(map[uint64]aggregator.RoundData)

	_ = k.Rounds.Walk(ctx, nil, func(id uint64, round aggregator.RoundData) (stop bool, err error) {
		rounds[id] = round
		return false, nil
	})

	return rounds
}
