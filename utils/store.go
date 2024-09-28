// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package utils

import (
	"cosmossdk.io/store/rootmulti"
	"cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKVStore retrieves the KVStore for the specified module from the context.
func GetKVStore(ctx sdk.Context, moduleName string) types.KVStore {
	return ctx.KVStore(ctx.MultiStore().(*rootmulti.Store).StoreKeysByName()[moduleName])
}
