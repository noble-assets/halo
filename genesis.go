// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package halo

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/types/entitlements"
)

func InitGenesis(ctx sdk.Context, k *keeper.Keeper, cdc address.Codec, genesis types.GenesisState) {
	if err := k.SetAggregatorOwner(ctx, genesis.AggregatorState.Owner); err != nil {
		panic(err)
	}
	if err := k.SetLastRoundId(ctx, genesis.AggregatorState.LastRoundId); err != nil {
		panic(err)
	}
	if err := k.SetNextPrice(ctx, genesis.AggregatorState.NextPrice); err != nil {
		panic(err)
	}
	for id, round := range genesis.AggregatorState.Rounds {
		if err := k.SetRound(ctx, id, round); err != nil {
			panic(err)
		}
	}

	if err := k.SetEntitlementsOwner(ctx, genesis.EntitlementsState.Owner); err != nil {
		panic(err)
	}
	for method, enabled := range genesis.EntitlementsState.PublicCapabilities {
		if err := k.SetPublicCapability(ctx, method, enabled); err != nil {
			panic(err)
		}
	}
	for _, entry := range genesis.EntitlementsState.RoleCapabilities {
		if err := k.SetRoleCapability(ctx, entry.Method, entry.Role, entry.Enabled); err != nil {
			panic(err)
		}
	}
	for _, entry := range genesis.EntitlementsState.UserRoles {
		user, _ := cdc.StringToBytes(entry.User)
		if err := k.SetUserRole(ctx, user, entry.Role, entry.Enabled); err != nil {
			panic(err)
		}
	}

	if err := k.SetOwner(ctx, genesis.Owner); err != nil {
		panic(err)
	}
	for account, nonce := range genesis.Nonces {
		address, _ := cdc.StringToBytes(account)
		if err := k.SetNonce(ctx, address, nonce); err != nil {
			panic(err)
		}
	}
}

func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		AggregatorState: aggregator.GenesisState{
			Owner:       k.GetAggregatorOwner(ctx),
			LastRoundId: k.GetLastRoundId(ctx),
			NextPrice:   k.GetNextPrice(ctx),
			Rounds:      k.GetRounds(ctx),
		},
		EntitlementsState: entitlements.GenesisState{
			Owner:              k.GetEntitlementsOwner(ctx),
			Paused:             k.GetPaused(ctx),
			PublicCapabilities: k.GetPublicCapabilities(ctx),
			RoleCapabilities:   k.GetAllCapabilityRoles(ctx),
			UserRoles:          k.GetAllUserRoles(ctx),
		},
		Owner:  k.GetOwner(ctx),
		Nonces: k.GetNonces(ctx),
	}
}
