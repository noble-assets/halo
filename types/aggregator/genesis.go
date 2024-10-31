// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package aggregator

import (
	"fmt"

	"cosmossdk.io/core/address"
)

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func (gs *GenesisState) Validate(cdc address.Codec) error {
	if gs.Reporter != "" {
		if _, err := cdc.StringToBytes(gs.Reporter); err != nil {
			return fmt.Errorf("invalid reporter address (%s): %s", gs.Reporter, err)
		}
	}

	return nil
}
