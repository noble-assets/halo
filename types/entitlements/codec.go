// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

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

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
