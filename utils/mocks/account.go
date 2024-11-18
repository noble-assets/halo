// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package mocks

import (
	"context"

	"github.com/noble-assets/halo/v3/types"

	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/codec"
)

var _ types.AccountKeeper = AccountKeeper{}

type AccountKeeper struct {
	Accounts map[string]sdk.AccountI
}

func (AccountKeeper) AddressCodec() address.Codec {
	return codec.NewBech32Codec("noble")
}

func (k AccountKeeper) GetAccount(_ context.Context, addr sdk.AccAddress) sdk.AccountI {
	// NOTE: The mock BankKeeper already sets the Bech32 prefix.
	return k.Accounts[addr.String()]
}
