package entitlements

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func (gs *GenesisState) Validate() error {
	if gs.Owner != "" {
		if _, err := sdk.AccAddressFromBech32(gs.Owner); err != nil {
			return fmt.Errorf("invalid entitlements owner address (%s): %s", gs.Owner, err)
		}
	}

	for _, entry := range gs.UserRoles {
		if _, err := sdk.AccAddressFromBech32(entry.User); err != nil {
			return fmt.Errorf("invalid entitlements user address (%s): %s", entry.User, err)
		}
	}

	return nil
}
