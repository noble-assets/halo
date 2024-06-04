package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type Account struct {
	Key     secp256k1.PrivKey
	Address string
	Invalid string
	Bytes   []byte
}

func TestAccount() Account {
	key := secp256k1.GenPrivKey()
	bytes := key.PubKey().Address().Bytes()
	address, _ := sdk.Bech32ifyAddressBytes("noble", bytes)
	invalid, _ := sdk.Bech32ifyAddressBytes("cosmos", bytes)

	return Account{
		Key:     key,
		Address: address,
		Invalid: invalid,
		Bytes:   bytes,
	}
}
