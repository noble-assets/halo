package mocks

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/noble-assets/halo/v2"
	"github.com/noble-assets/halo/v2/keeper"
	"github.com/noble-assets/halo/v2/types"
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

func HaloKeeperWithKeepers(t testing.TB, account AccountKeeper, bank BankKeeper) (*keeper.Keeper, sdk.Context) {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	tkey := storetypes.NewTransientStoreKey("transient_authority")
	wrapper := testutil.DefaultContextWithDB(t, key, tkey)

	cfg := MakeTestEncodingConfig("noble")
	types.RegisterInterfaces(cfg.InterfaceRegistry)

	cdc := address.NewBech32Codec("noble")
	k := keeper.NewKeeper(
		cfg.Codec,
		runtime.NewKVStoreService(key),
		runtime.ProvideEventService(),
		runtime.ProvideHeaderInfoService(&runtime.AppBuilder{}),
		"uusyc",
		"uusdc",
		cdc,
		account,
		bank,
		cfg.InterfaceRegistry,
	)

	bank = bank.WithSendCoinsRestriction(k.SendRestrictionFn)
	k.SetBankKeeper(bank)

	halo.InitGenesis(wrapper.Ctx, k, cdc, *types.DefaultGenesisState())
	return k, wrapper.Ctx
}

// MakeTestEncodingConfig is a modified testutil.MakeTestEncodingConfig that
// sets a custom Bech32 prefix in the interface registry.
func MakeTestEncodingConfig(prefix string, modules ...module.AppModuleBasic) moduletestutil.TestEncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectestutil.CodecOptions{
		AccAddressPrefix: prefix,
	}.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	encCfg := moduletestutil.TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             aminoCodec,
	}

	mb := module.NewBasicManager(modules...)

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	mb.RegisterLegacyAminoCodec(encCfg.Amino)
	mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}
