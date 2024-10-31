// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/noble-assets/halo/v2/types/aggregator"
	"github.com/spf13/cobra"
)

func GetAggregatorQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "aggregator",
		Short:                      "Querying commands for the aggregator submodule",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(QueryReporter())

	return cmd
}

func QueryReporter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reporter",
		Short: "Query the aggregator reporter",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := aggregator.NewQueryClient(clientCtx)

			res, err := queryClient.Reporter(context.Background(), &aggregator.QueryReporter{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
