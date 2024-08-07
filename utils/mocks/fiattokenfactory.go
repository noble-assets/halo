package mocks

import (
	ftftypes "github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types"
)

var _ types.FiatTokenFactoryKeeper = FTFKeeper{}

type FTFKeeper struct {
	Paused bool
}

func (k FTFKeeper) GetBlacklisted(_ sdk.Context, bz []byte) (ftftypes.Blacklisted, bool) {
	return ftftypes.Blacklisted{AddressBz: bz}, false // TODO
}

func (k FTFKeeper) GetPaused(_ sdk.Context) ftftypes.Paused {
	return ftftypes.Paused{Paused: k.Paused}
}

func (k FTFKeeper) GetMintingDenom(_ sdk.Context) ftftypes.MintingDenom {
	return ftftypes.MintingDenom{Denom: "uusdc"}
}