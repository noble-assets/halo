package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		AggregatorState:   aggregator.DefaultGenesisState(),
		EntitlementsState: entitlements.DefaultGenesisState(),
	}
}

func (gs *GenesisState) Validate() error {
	if err := gs.AggregatorState.Validate(); err != nil {
		return err
	}

	if err := gs.EntitlementsState.Validate(); err != nil {
		return err
	}

	if gs.Owner != "" {
		if _, err := sdk.AccAddressFromBech32(gs.Owner); err != nil {
			return fmt.Errorf("invalid owner address (%s): %s", gs.Owner, err)
		}
	}

	for address := range gs.Nonces {
		if _, err := sdk.AccAddressFromBech32(address); err != nil {
			return fmt.Errorf("invalid user address (%s): %s", address, err)
		}
	}

	return nil
}
