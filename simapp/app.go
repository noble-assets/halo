// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package simapp

import (
	_ "embed"
	"io"
	"os"
	"path/filepath"

	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	_ "github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	_ "github.com/cosmos/cosmos-sdk/x/bank"
	_ "github.com/cosmos/cosmos-sdk/x/consensus"
	_ "github.com/cosmos/cosmos-sdk/x/staking"
	_ "github.com/noble-assets/halo/v3"

	// Cosmos Modules
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	// Custom Modules
	halokeeper "github.com/noble-assets/halo/v3/keeper"
)

var DefaultNodeHome string

//go:embed app.yaml
var AppConfigYAML []byte

var (
	_ runtime.AppI            = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// Cosmos Modules
	AccountKeeper   authkeeper.AccountKeeper
	BankKeeper      bankkeeper.Keeper
	ConsensusKeeper consensuskeeper.Keeper
	StakingKeeper   *stakingkeeper.Keeper
	// Custom Modules
	HaloKeeper *halokeeper.Keeper
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
}

// AppConfig returns the default app config.
func AppConfig() depinject.Config {
	return depinject.Configs(
		appconfig.LoadYAML(AppConfigYAML),
		depinject.Supply(
			// supply custom module basics
			map[string]module.AppModuleBasic{
				genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			},
		),
	)
}

// NewSimApp returns a reference to an initialized SimApp.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) (*SimApp, error) {
	var (
		app        = &SimApp{}
		appBuilder *runtime.AppBuilder
	)

	if err := depinject.Inject(
		depinject.Configs(
			AppConfig(),
			depinject.Supply(
				logger,
				appOpts,
			),
		),
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		// Cosmos Modules
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.ConsensusKeeper,
		&app.StakingKeeper,
		// Custom Modules
		&app.HaloKeeper,
	); err != nil {
		return nil, err
	}

	app.App = appBuilder.Build(db, traceStore, baseAppOptions...)

	if err := app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		return nil, err
	}

	if err := app.Load(loadLatest); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

func (app *SimApp) SimulationManager() *module.SimulationManager {
	return nil
}

func (app *SimApp) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}
