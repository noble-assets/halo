package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

func DefaultGenesisState() *GenesisState {
	// To avoid an import cycle, default role capabilities are set here.
	var roleCapabilities []entitlements.RoleCapability
	for i := entitlements.ROLE_DOMESTIC_FEEDER; i <= entitlements.ROLE_INTERNATIONAL_SDYF; i++ {
		roleCapabilities = append(roleCapabilities, entitlements.RoleCapability{
			Method:  sdk.MsgTypeURL(&MsgBurn{}),
			Role:    i,
			Enabled: true,
		})
		roleCapabilities = append(roleCapabilities, entitlements.RoleCapability{
			Method:  sdk.MsgTypeURL(&MsgDeposit{}),
			Role:    i,
			Enabled: true,
		})
		roleCapabilities = append(roleCapabilities, entitlements.RoleCapability{
			Method:  sdk.MsgTypeURL(&MsgDepositFor{}),
			Role:    i,
			Enabled: true,
		})
		roleCapabilities = append(roleCapabilities, entitlements.RoleCapability{
			Method:  sdk.MsgTypeURL(&MsgWithdraw{}),
			Role:    i,
			Enabled: true,
		})
		roleCapabilities = append(roleCapabilities, entitlements.RoleCapability{
			Method:  sdk.MsgTypeURL(&MsgWithdrawTo{}),
			Role:    i,
			Enabled: true,
		})
		roleCapabilities = append(roleCapabilities, entitlements.RoleCapability{
			Method:  "transfer",
			Role:    i,
			Enabled: true,
		})
	}

	entitlementsState := entitlements.DefaultGenesisState()
	entitlementsState.RoleCapabilities = append(entitlementsState.RoleCapabilities, roleCapabilities...)

	return &GenesisState{
		AggregatorState:   aggregator.DefaultGenesisState(),
		EntitlementsState: entitlementsState,
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
