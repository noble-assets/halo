// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package cli

import (
	"errors"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/spf13/cobra"
)

func GetAggregatorTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "aggregator",
		Short:                      "Transactions commands for the aggregator submodule",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(TxTransmit())
	cmd.AddCommand(TxSetNextPrice())
	cmd.AddCommand(TxAggregatorTransferOwnership())

	return cmd
}

func TxTransmit() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "transmit [answer] [updated-at]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			answer, ok := math.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid answer")
			}

			updatedAt, err := strconv.Atoi(args[1])
			if err != nil {
				return errors.New("invalid updated-at")
			}

			msg := &aggregator.MsgTransmit{
				Signer:    clientCtx.GetFromAddress().String(),
				Answer:    answer,
				UpdatedAt: uint32(updatedAt),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxSetNextPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-next-price [next-price]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nextPrice, ok := math.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid next price")
			}

			msg := &aggregator.MsgSetNextPrice{
				Signer:    clientCtx.GetFromAddress().String(),
				NextPrice: nextPrice,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxAggregatorTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [new-reporter]",
		Short: "Transfer ownership of submodule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &aggregator.MsgTransferOwnership{
				Signer:      clientCtx.GetFromAddress().String(),
				NewReporter: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
