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
