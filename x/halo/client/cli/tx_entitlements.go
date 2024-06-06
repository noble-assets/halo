package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/spf13/cobra"
)

func GetEntitlementsTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "entitlements",
		Short:                      "Transactions commands for the entitlements submodule",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// TODO: Add missing commands.
	cmd.AddCommand(TxSetUserRole())

	cmd.AddCommand(TxPause())
	cmd.AddCommand(TxUnpause())
	cmd.AddCommand(TxEntitlementsTransferOwnership())

	return cmd
}

func TxSetUserRole() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-user-role [user] [role] [enabled]",
		Short: "Transaction that pauses the submodule",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var role entitlements.Role
			switch args[1] {
			case "fund-admin":
				role = entitlements.ROLE_FUND_ADMIN
			// TODO: More roles
			default:
				return fmt.Errorf("unknown role: %s", args[1])
			}

			enabled, _ := strconv.ParseBool(args[2])

			msg := &entitlements.MsgSetUserRole{
				Signer:  clientCtx.GetFromAddress().String(),
				User:    args[0],
				Role:    role,
				Enabled: enabled,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxPause() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause",
		Short: "Transaction that pauses the submodule",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &entitlements.MsgPause{
				Signer: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxUnpause() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpause",
		Short: "Transaction that unpauses the submodule",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &entitlements.MsgUnpause{
				Signer: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxEntitlementsTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [new-owner]",
		Short: "Transfer ownership of submodule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &entitlements.MsgTransferOwnership{
				Signer:   clientCtx.GetFromAddress().String(),
				NewOwner: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
