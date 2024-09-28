// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package utils

import (
	"github.com/cometbft/cometbft/crypto/secp256k1"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmos "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Account struct {
	Key     secp256k1.PrivKey
	PubKey  *codectypes.Any
	Address string
	Invalid string
	Bytes   []byte
}

func TestAccount() Account {
	key := secp256k1.GenPrivKey()
	pubKey, _ := codectypes.NewAnyWithValue(&cosmos.PubKey{
		Key: key.PubKey().Bytes(),
	})
	bytes := key.PubKey().Address().Bytes()
	address, _ := sdk.Bech32ifyAddressBytes("noble", bytes)
	invalid, _ := sdk.Bech32ifyAddressBytes("cosmos", bytes)

	return Account{
		Key:     key,
		PubKey:  pubKey,
		Address: address,
		Invalid: invalid,
		Bytes:   bytes,
	}
}
