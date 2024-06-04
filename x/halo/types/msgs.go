package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

//

var _ legacytx.LegacyMsg = &MsgDeposit{}

func (msg *MsgDeposit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	return nil
}

func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgDeposit) Route() string { return ModuleName }

func (*MsgDeposit) Type() string { return "halo/Deposit" }

//

var _ legacytx.LegacyMsg = &MsgDepositFor{}

func (msg *MsgDepositFor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return fmt.Errorf("invalid recipient address (%s): %w", msg.Recipient, err)
	}

	return nil
}

func (msg *MsgDepositFor) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgDepositFor) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgDepositFor) Route() string { return ModuleName }

func (*MsgDepositFor) Type() string { return "halo/DepositFor" }

//

var _ legacytx.LegacyMsg = &MsgWithdraw{}

func (msg *MsgWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	return nil
}

func (msg *MsgWithdraw) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgWithdraw) Route() string { return ModuleName }

func (*MsgWithdraw) Type() string { return "halo/Withdraw" }

//

var _ legacytx.LegacyMsg = &MsgWithdrawTo{}

func (msg *MsgWithdrawTo) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return fmt.Errorf("invalid recipient address (%s): %w", msg.Recipient, err)
	}

	return nil
}

func (msg *MsgWithdrawTo) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgWithdrawTo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgWithdrawTo) Route() string { return ModuleName }

func (*MsgWithdrawTo) Type() string { return "halo/WithdrawTo" }

//

var _ legacytx.LegacyMsg = &MsgWithdrawToAdmin{}

func (msg *MsgWithdrawToAdmin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.From); err != nil {
		return fmt.Errorf("invalid from address (%s): %w", msg.From, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return fmt.Errorf("invalid recipient address (%s): %w", msg.Recipient, err)
	}

	return nil
}

func (msg *MsgWithdrawToAdmin) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgWithdrawToAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgWithdrawToAdmin) Route() string { return ModuleName }

func (*MsgWithdrawToAdmin) Type() string { return "halo/WithdrawToAdmin" }

//

var _ legacytx.LegacyMsg = &MsgBurn{}

func (msg *MsgBurn) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	return nil
}

func (msg *MsgBurn) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgBurn) Route() string { return ModuleName }

func (*MsgBurn) Type() string { return "halo/Burn" }

//

var _ legacytx.LegacyMsg = &MsgBurnFor{}

func (msg *MsgBurnFor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.From); err != nil {
		return fmt.Errorf("invalid from address (%s): %w", msg.From, err)
	}

	return nil
}

func (msg *MsgBurnFor) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgBurnFor) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgBurnFor) Route() string { return ModuleName }

func (*MsgBurnFor) Type() string { return "halo/BurnFor" }

//

var _ legacytx.LegacyMsg = &MsgMint{}

func (msg *MsgMint) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.To); err != nil {
		return fmt.Errorf("invalid to address (%s): %w", msg.To, err)
	}

	return nil
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgMint) Route() string { return ModuleName }

func (*MsgMint) Type() string { return "halo/Mint" }

//

var _ legacytx.LegacyMsg = &MsgTradeToFiat{}

func (msg *MsgTradeToFiat) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return fmt.Errorf("invalid recipient address (%s): %w", msg.Recipient, err)
	}

	return nil
}

func (msg *MsgTradeToFiat) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgTradeToFiat) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgTradeToFiat) Route() string { return ModuleName }

func (*MsgTradeToFiat) Type() string { return "halo/TradeToFiat" }

//

var _ legacytx.LegacyMsg = &MsgTransferOwnership{}

func (msg *MsgTransferOwnership) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return fmt.Errorf("invalid signer address (%s): %w", msg.Signer, err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.NewOwner); err != nil {
		return fmt.Errorf("invalid new owner address (%s): %w", msg.NewOwner, err)
	}

	return nil
}

func (msg *MsgTransferOwnership) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}

func (msg *MsgTransferOwnership) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (*MsgTransferOwnership) Route() string { return ModuleName }

func (*MsgTransferOwnership) Type() string { return "halo/TransferOwnership" }
