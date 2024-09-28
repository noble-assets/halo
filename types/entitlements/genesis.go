package entitlements

import (
	"fmt"

	"cosmossdk.io/core/address"
)

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func (gs *GenesisState) Validate(cdc address.Codec) error {
	if gs.Owner != "" {
		if _, err := cdc.StringToBytes(gs.Owner); err != nil {
			return fmt.Errorf("invalid entitlements owner address (%s): %s", gs.Owner, err)
		}
	}

	for _, entry := range gs.UserRoles {
		if _, err := cdc.StringToBytes(entry.User); err != nil {
			return fmt.Errorf("invalid entitlements user address (%s): %s", entry.User, err)
		}
	}

	return nil
}
