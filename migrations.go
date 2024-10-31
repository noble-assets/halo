// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package halo

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types/aggregator"
	aggregatorv1 "github.com/noble-assets/halo/v2/types/aggregator/v1"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *keeper.Keeper

	v1Rounds collections.Map[uint64, aggregatorv1.RoundData]
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper *keeper.Keeper, service store.KVStoreService, cdc codec.Codec) Migrator {
	builder := collections.NewSchemaBuilder(service)

	migrator := Migrator{
		keeper: keeper,

		v1Rounds: collections.NewMap(builder, aggregator.RoundPrefix, "aggregator_rounds", collections.Uint64Key, codec.CollValue[aggregatorv1.RoundData](cdc)),
	}

	if _, err := builder.Build(); err != nil {
		panic(err)
	}

	return migrator
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	rounds := make(map[uint64]aggregator.RoundData)
	err := m.v1Rounds.Walk(ctx, nil, func(id uint64, round aggregatorv1.RoundData) (stop bool, err error) {
		rounds[id] = aggregator.RoundData{
			Answer:    round.Answer,
			UpdatedAt: uint32(round.UpdatedAt),
		}

		return false, nil
	})
	if err != nil {
		return err
	}

	for id, round := range rounds {
		if err = m.keeper.Rounds.Set(ctx, id, round); err != nil {
			return err
		}
	}

	return nil
}
