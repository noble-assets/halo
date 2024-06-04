package mocks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/noble-assets/halo/x/halo/types"
)

var _ types.AccountKeeper = AccountKeeper{}

type AccountKeeper struct {
	Accounts map[string]authtypes.AccountI
}

func (k AccountKeeper) GetAccount(_ sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	// NOTE: The mock BankKeeper already sets the Bech32 prefix.
	return k.Accounts[addr.String()]
}
