// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/noble-assets/halo/v2/types/entitlements"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	aggregator.RegisterLegacyAminoCodec(cdc)
	entitlements.RegisterLegacyAminoCodec(cdc)

	cdc.RegisterConcrete(&MsgDeposit{}, "halo/Deposit", nil)
	cdc.RegisterConcrete(&MsgDepositFor{}, "halo/DepositFor", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "halo/Withdraw", nil)
	cdc.RegisterConcrete(&MsgWithdrawTo{}, "halo/WithdrawTo", nil)
	cdc.RegisterConcrete(&MsgWithdrawToAdmin{}, "halo/WithdrawToAdmin", nil)

	cdc.RegisterConcrete(&MsgBurn{}, "halo/Burn", nil)
	cdc.RegisterConcrete(&MsgBurnFor{}, "halo/BurnFor", nil)
	cdc.RegisterConcrete(&MsgMint{}, "halo/Mint", nil)
	cdc.RegisterConcrete(&MsgTradeToFiat{}, "halo/TradeToFiat", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "halo/TransferOwnership", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	aggregator.RegisterInterfaces(registry)
	entitlements.RegisterInterfaces(registry)

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgDeposit{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgDepositFor{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgWithdraw{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgWithdrawTo{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgWithdrawToAdmin{})

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgBurn{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgBurnFor{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgMint{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTradeToFiat{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransferOwnership{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
