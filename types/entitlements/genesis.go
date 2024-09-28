// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

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
