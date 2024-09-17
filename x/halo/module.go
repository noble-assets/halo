package halo

import (
	"context"
	"encoding/json"
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	aggregatorv1 "github.com/noble-assets/halo/api/aggregator/v1"
	entitlementsv1 "github.com/noble-assets/halo/api/entitlements/v1"
	modulev1 "github.com/noble-assets/halo/api/module/v1"
	halov1 "github.com/noble-assets/halo/api/v1"
	"github.com/noble-assets/halo/x/halo/keeper"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
)

// ConsensusVersion defines the current x/halo module consensus version.
const ConsensusVersion = 2

var (
	_ module.AppModuleBasic      = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasGenesisBasics    = AppModuleBasic{}
	_ module.HasServices         = AppModule{}
)

//

type AppModuleBasic struct{}

func NewAppModuleBasic() AppModuleBasic {
	return AppModuleBasic{}
}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(reg codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}

	if err := aggregator.RegisterQueryHandlerClient(context.Background(), mux, aggregator.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}

	if err := entitlements.RegisterQueryHandlerClient(context.Background(), mux, entitlements.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var genesis types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genesis); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return genesis.Validate()
}

//

type AppModule struct {
	AppModuleBasic

	keeper *keeper.Keeper
}

func NewAppModule(keeper *keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(),
		keeper:         keeper,
	}
}

func (AppModule) IsOnePerModuleType() {}

func (AppModule) IsAppModule() {}

func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

func (m AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
	var genesis types.GenesisState
	cdc.MustUnmarshalJSON(bz, &genesis)

	InitGenesis(ctx, m.keeper, genesis)
}

func (m AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesis := ExportGenesis(ctx, m.keeper)
	return cdc.MustMarshalJSON(genesis)
}

func (m AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(m.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(m.keeper))

	aggregator.RegisterMsgServer(cfg.MsgServer(), keeper.NewAggregatorMsgServer(m.keeper))
	aggregator.RegisterQueryServer(cfg.QueryServer(), keeper.NewAggregatorQueryServer(m.keeper))

	entitlements.RegisterMsgServer(cfg.MsgServer(), keeper.NewEntitlementsMsgServer(m.keeper))
	entitlements.RegisterQueryServer(cfg.QueryServer(), keeper.NewEntitlementsQueryServer(m.keeper))
}

//

func (AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: halov1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "Deposit",
					Use:            "deposit [amount]",
					Short:          "Deposit a specific amount of underlying assets for USYC",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
				{
					RpcMethod: "DepositFor",
					Use:       "deposit-for [recipient] [amount]",
					Short:     "Deposit a specific amount of underlying assets for USYC, with a recipient",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "recipient"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "Withdraw",
					Use:       "withdraw [amount] [signature]",
					Short:     "Withdraw a specific amount of USYC for underlying assets",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "amount"},
						{ProtoField: "signature"},
					},
				},
				{
					RpcMethod: "WithdrawTo",
					Use:       "withdraw-to [recipient] [amount] [signature]",
					Short:     "Withdraw a specific amount of USYC for underlying assets, with a recipient",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "recipient"},
						{ProtoField: "amount"},
						{ProtoField: "signature"},
					},
				},
				{
					RpcMethod: "WithdrawToAdmin",
					Use:       "withdraw-to-admin [from] [recipient] [amount]",
					Short:     "Withdraw a specific amount of USYC as a fund admin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "from"},
						{ProtoField: "recipient"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod:      "Burn",
					Use:            "burn [amount]",
					Short:          "Transaction that burns tokens",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
				{
					RpcMethod: "BurnFor",
					Use:       "burn-for [from] [amount]",
					Short:     "Transaction that burns tokens from a specific account",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "from"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "Mint",
					Use:       "mint [to] [amount]",
					Short:     "Transaction that mints tokens",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "to"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "TradeToFiat",
					Use:       "trade-to-fiat [amount] [recipient]",
					Short:     "Withdraw underlying assets from module as a liquidity provider",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "amount"},
						{ProtoField: "recipient"},
					},
				},
				{
					RpcMethod:      "TransferOwnership",
					Use:            "transfer-ownership [new-owner]",
					Short:          "Transfer ownership of submodule",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "new_owner"}},
				},
			},
			SubCommands: map[string]*autocliv1.ServiceCommandDescriptor{
				"aggregator": {
					Service: aggregatorv1.Msg_ServiceDesc.ServiceName,
					RpcCommandOptions: []*autocliv1.RpcCommandOptions{
						{
							RpcMethod: "ReportBalance",
							Use:       "report-balance [principal] [interest] [total-supply] [next-price]",
							Short:     "Report the latest round with a new next price",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{
								{ProtoField: "principal"},
								{ProtoField: "interest"},
								{ProtoField: "total_supply"},
								{ProtoField: "next_price"},
							},
						},
						{
							RpcMethod:      "SetNextPrice",
							Use:            "set-next-price [next-price]",
							Short:          "Update the next price",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "next_price"}},
						},
						{
							RpcMethod:      "TransferOwnership",
							Use:            "transfer-ownership [new-owner]",
							Short:          "Transfer ownership of submodule",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "new_owner"}},
						},
					},
					Short: "Transactions commands for the aggregator submodule",
				},
				"entitlements": {
					Service: entitlementsv1.Msg_ServiceDesc.ServiceName,
					RpcCommandOptions: []*autocliv1.RpcCommandOptions{
						{
							RpcMethod: "SetPublicCapability",
							Use:       "set-public-capability [method] [enabled]",
							Short:     "Enable or disable a specific public method",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{
								{ProtoField: "method"},
								{ProtoField: "enabled"},
							},
						},
						{
							RpcMethod: "SetRoleCapability",
							Use:       "set-role-capability [role] [method] [enabled]",
							Short:     "Enable or disable a specific method for a role",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{
								{ProtoField: "role"},
								{ProtoField: "method"},
								{ProtoField: "enabled"},
							},
						},
						{
							RpcMethod: "SetUserRole",
							Use:       "set-user-role [user] [role] [enabled]",
							Short:     "Enable or disable a specific role for a user",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{
								{ProtoField: "user"},
								{ProtoField: "role"},
								{ProtoField: "enabled"},
							},
						},
						{
							RpcMethod: "Pause",
							Use:       "pause",
							Short:     "Transaction that pauses the submodule",
						},
						{
							RpcMethod: "Unpause",
							Use:       "unpause",
							Short:     "Transaction that unpauses the submodule",
						},
						{
							RpcMethod:      "TransferOwnership",
							Use:            "transfer-ownership [new-owner]",
							Short:          "Transfer ownership of submodule",
							PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "new_owner"}},
						},
					},
					Short: "Transactions commands for the entitlements submodule",
				},
			},
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:           halov1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
			SubCommands: map[string]*autocliv1.ServiceCommandDescriptor{
				"aggregator": {
					Service: aggregatorv1.Query_ServiceDesc.ServiceName,
					RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
				},
				"entitlements": {
					Service: entitlementsv1.Query_ServiceDesc.ServiceName,
					RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
				},
			},
		},
	}
}

//

func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Config            *modulev1.Module
	Cdc               codec.Codec
	StoreService      store.KVStoreService
	InterfaceRegistry codectypes.InterfaceRegistry

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
}

type ModuleOutputs struct {
	depinject.Out

	Keeper       *keeper.Keeper
	Module       appmodule.AppModule
	Restrictions banktypes.SendRestrictionFn
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Config.Denom,
		in.Config.Underlying,
		in.AccountKeeper,
		in.BankKeeper,
		in.InterfaceRegistry,
	)
	m := NewAppModule(k)

	return ModuleOutputs{Keeper: k, Module: m, Restrictions: k.SendRestrictionFn}
}
