// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/gogoproto/proto"
	"github.com/noble-assets/halo/v3/types/aggregator"
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

	cmd.AddCommand(QueryAggregatorOwner())
	cmd.AddCommand(QueryNextPrice())
	cmd.AddCommand(QueryRoundData())
	cmd.AddCommand(QueryRoundDetails())

	return cmd
}

func QueryAggregatorOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner",
		Short: "Query the submodule's owner",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := aggregator.NewQueryClient(clientCtx)

			res, err := queryClient.Owner(context.Background(), &aggregator.QueryOwner{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func QueryNextPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-price",
		Short: "Query the next price",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := aggregator.NewQueryClient(clientCtx)

			res, err := queryClient.NextPrice(context.Background(), &aggregator.QueryNextPrice{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func QueryRoundData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "round-data (id)",
		Short: "Query the data of a the latest or a specific round",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := aggregator.NewQueryClient(clientCtx)

			var res proto.Message
			var err error

			if len(args) == 1 {
				id, parseErr := strconv.ParseUint(args[0], 10, 64)
				if parseErr != nil {
					return parseErr
				}

				res, err = queryClient.RoundData(context.Background(), &aggregator.QueryRoundData{
					RoundId: id,
				})
			} else {
				res, err = queryClient.LatestRoundData(context.Background(), &aggregator.QueryLatestRoundData{})
			}

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func QueryRoundDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "round-details (id)",
		Short: "Query the details of a the latest or a specific round",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := aggregator.NewQueryClient(clientCtx)

			var res proto.Message
			var err error

			if len(args) == 1 {
				id, parseErr := strconv.ParseUint(args[0], 10, 64)
				if parseErr != nil {
					return parseErr
				}

				res, err = queryClient.RoundDetails(context.Background(), &aggregator.QueryRoundDetails{
					RoundId: id,
				})
			} else {
				res, err = queryClient.LatestRoundDetails(context.Background(), &aggregator.QueryLatestRoundDetails{})
			}

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
