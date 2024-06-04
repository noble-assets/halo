package halo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genesis types.GenesisState) {
	k.SetAggregatorOwner(ctx, genesis.AggregatorState.Owner)
	k.SetLastRoundId(ctx, genesis.AggregatorState.LastRoundId)
	k.SetNextPrice(ctx, genesis.AggregatorState.NextPrice)
	for id, round := range genesis.AggregatorState.Rounds {
		k.SetRound(ctx, id, round)
	}

	k.SetEntitlementsOwner(ctx, genesis.EntitlementsState.Owner)
	for method, enabled := range genesis.EntitlementsState.PublicCapabilities {
		k.SetPublicCapability(ctx, method, enabled)
	}
	for _, entry := range genesis.EntitlementsState.RoleCapabilities {
		k.SetRoleCapability(ctx, entry.Method, entry.Role, entry.Enabled)
	}
	for _, entry := range genesis.EntitlementsState.UserRoles {
		user := sdk.MustAccAddressFromBech32(entry.User)
		k.SetUserRole(ctx, user, entry.Role, entry.Enabled)
	}

	k.SetOwner(ctx, genesis.Owner)
	for account, nonce := range genesis.Nonces {
		address := sdk.MustAccAddressFromBech32(account)
		k.SetNonce(ctx, address, nonce)
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
			Owner:  k.GetEntitlementsOwner(ctx),
			Paused: k.GetPaused(ctx),
			// TODO
		},
		Owner:  k.GetOwner(ctx),
		Nonces: k.GetNonces(ctx),
	}
}
