package aggregator

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
			return fmt.Errorf("invalid aggregator owner address (%s): %s", gs.Owner, err)
		}
	}

	return nil
}
