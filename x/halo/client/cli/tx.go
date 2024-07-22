package cli

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transactions commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetAggregatorTxCmd())
	cmd.AddCommand(GetEntitlementsTxCmd())

	cmd.AddCommand(TxDeposit())
	cmd.AddCommand(TxDepositFor())
	cmd.AddCommand(TxWithdraw())
	cmd.AddCommand(TxWithdrawTo())
	cmd.AddCommand(TxWithdrawToAdmin())
	cmd.AddCommand(TxBurn())
	cmd.AddCommand(TxBurnFor())
	cmd.AddCommand(TxMint())
	cmd.AddCommand(TxTradeToFiat())
	cmd.AddCommand(TxTransferOwnership())

	return cmd
}

func TxDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [amount]",
		Short: "Deposit a specific amount of underlying assets for USYC",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgDeposit{
				Signer: clientCtx.GetFromAddress().String(),
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxDepositFor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-for [recipient] [amount]",
		Short: "Deposit a specific amount of underlying assets for USYC, with a recipient",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgDepositFor{
				Signer:    clientCtx.GetFromAddress().String(),
				Recipient: args[0],
				Amount:    amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [amount] [signature]",
		Short: "Withdraw a specific amount of USYC for underlying assets",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid amount")
			}

			signature, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				return err
			}

			msg := &types.MsgWithdraw{
				Signer:    clientCtx.GetFromAddress().String(),
				Amount:    amount,
				Signature: signature,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxWithdrawTo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-to [recipient] [amount] [signature]",
		Short: "Withdraw a specific amount of USYC for underlying assets, with a recipient",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return errors.New("invalid amount")
			}

			signature, err := base64.StdEncoding.DecodeString(args[2])
			if err != nil {
				return err
			}

			msg := &types.MsgWithdrawTo{
				Signer:    clientCtx.GetFromAddress().String(),
				Recipient: args[0],
				Amount:    amount,
				Signature: signature,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxWithdrawToAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-to-admin [from] [recipient] [amount]",
		Short: "Withdraw a specific amount of USYC as a fund admin",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgWithdrawToAdmin{
				Signer:    clientCtx.GetFromAddress().String(),
				From:      args[0],
				Recipient: args[1],
				Amount:    amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount]",
		Short: "Transaction that burns tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgBurn{
				Signer: clientCtx.GetFromAddress().String(),
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxBurnFor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-for [from] [amount]",
		Short: "Transaction that burns tokens from a specific account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgBurnFor{
				Signer: clientCtx.GetFromAddress().String(),
				From:   args[0],
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [to] [amount]",
		Short: "Transaction that mints tokens",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgMint{
				Signer: clientCtx.GetFromAddress().String(),
				To:     args[0],
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxTradeToFiat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trade-to-fiat [amount] [recipient]",
		Short: "Withdraw underlying assets from module as a liquidity provider",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgTradeToFiat{
				Signer:    clientCtx.GetFromAddress().String(),
				Amount:    amount,
				Recipient: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [new-owner]",
		Short: "Transfer ownership of module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgTransferOwnership{
				Signer:   clientCtx.GetFromAddress().String(),
				NewOwner: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
