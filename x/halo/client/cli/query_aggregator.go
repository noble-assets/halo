package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gogo/protobuf/proto"
	"github.com/noble-assets/halo/x/halo/types/aggregator"
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
