// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package cli

import (
	"errors"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/noble-assets/halo/v3/types/aggregator"
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

	cmd.AddCommand(TxReportBalance())
	cmd.AddCommand(TxSetNextPrice())
	cmd.AddCommand(TxAggregatorTransferOwnership())

	return cmd
}

func TxReportBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-balance [principal] [interest] [total-supply] [next-price]",
		Short: "Transfer ownership of submodule",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			principal, ok := math.NewIntFromString(args[0])
			if !ok {
				return errors.New("invalid principal")
			}

			interest, ok := math.NewIntFromString(args[1])
			if !ok {
				return errors.New("invalid interest")
			}

			totalSupply, ok := math.NewIntFromString(args[2])
			if !ok {
				return errors.New("invalid total supply")
			}

			nextPrice, ok := math.NewIntFromString(args[3])
			if !ok {
				return errors.New("invalid next price")
			}

			msg := &aggregator.MsgReportBalance{
				Signer:      clientCtx.GetFromAddress().String(),
				Principal:   principal,
				Interest:    interest,
				TotalSupply: totalSupply,
				NextPrice:   nextPrice,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxSetNextPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-next-price [next-price]",
		Short: "Transfer ownership of submodule",
		Args:  cobra.ExactArgs(1),
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
		Use:   "transfer-ownership [new-owner]",
		Short: "Transfer ownership of submodule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &aggregator.MsgTransferOwnership{
				Signer:   clientCtx.GetFromAddress().String(),
				NewOwner: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
