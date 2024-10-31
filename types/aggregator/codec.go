// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package aggregator

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgTransmit{}, "halo/aggregator/Transmit", nil)
	cdc.RegisterConcrete(&MsgSetNextPrice{}, "halo/aggregator/SetNextPrice", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "halo/aggregator/TransferOwnership", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransmit{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSetNextPrice{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransferOwnership{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
