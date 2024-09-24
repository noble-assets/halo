package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/gogoproto/proto"
	"github.com/noble-assets/halo/v2/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetAggregatorQueryCmd())
	cmd.AddCommand(GetEntitlementsQueryCmd())

	cmd.AddCommand(QueryOwner())
	cmd.AddCommand(QueryNonces())

	return cmd
}

func QueryOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner",
		Short: "Query the module's owner",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Owner(context.Background(), &types.QueryOwner{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func QueryNonces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nonces (address)",
		Short: "Query the all or a specific withdrawal nonce",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			var res proto.Message
			var err error

			if len(args) == 1 {
				res, err = queryClient.Nonce(context.Background(), &types.QueryNonce{
					Address: args[0],
				})
			} else {
				pagination, parseErr := client.ReadPageRequest(cmd.Flags())
				if parseErr != nil {
					return parseErr
				}

				res, err = queryClient.Nonces(context.Background(), &types.QueryNonces{
					Pagination: pagination,
				})
			}

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nonces")

	return cmd
}
