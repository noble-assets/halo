package mocks

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
)

func HaloKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	return HaloKeeperWithKeepers(
		t,
		AccountKeeper{},
		BankKeeper{
			Restriction: NoOpSendRestrictionFn,
		},
	)
}

func HaloKeeperWithKeepers(_ testing.TB, account AccountKeeper, bank BankKeeper) (*keeper.Keeper, sdk.Context) {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	tkey := storetypes.NewTransientStoreKey("transient_halo")
	ctx := testutil.DefaultContext(key, tkey)

	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		key,
		"uusyc",
		"uusdc",
		account,
		nil,
		registry,
	)

	bank = bank.WithSendCoinsRestriction(bank.Restriction)
	k.SetBankKeeper(bank)

	halo.InitGenesis(ctx, k, *types.DefaultGenesisState())

	return k, ctx
}
