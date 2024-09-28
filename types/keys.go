// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

import authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

const ModuleName = "halo"

var ModuleAddress = authtypes.NewModuleAddress(ModuleName)

var (
	OwnerKey    = []byte("owner")
	NoncePrefix = []byte("nonce/")
)

func NonceKey(address []byte) []byte {
	return append(NoncePrefix, address...)
}
