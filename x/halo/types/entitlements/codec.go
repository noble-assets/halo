package entitlements

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetPublicCapability{}, "halo/entitlements/SetPublicCapability", nil)
	cdc.RegisterConcrete(&MsgSetRoleCapability{}, "halo/entitlements/SetRoleCapability", nil)
	cdc.RegisterConcrete(&MsgSetUserRole{}, "halo/entitlements/SetUserRole", nil)

	cdc.RegisterConcrete(&MsgPause{}, "halo/entitlements/Pause", nil)
	cdc.RegisterConcrete(&MsgUnpause{}, "halo/entitlements/Unpause", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "halo/entitlements/TransferOwnership", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSetPublicCapability{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSetRoleCapability{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSetUserRole{})

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgPause{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUnpause{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransferOwnership{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
