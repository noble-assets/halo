package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/noble-assets/halo/x/halo/types/entitlements"
	"github.com/spf13/cobra"
)

func GetEntitlementsQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "entitlements",
		Short:                      "Querying commands for the entitlements submodule",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(QueryEntitlementsOwner())
	cmd.AddCommand(QueryPaused())
	cmd.AddCommand(QueryPublicCapabilities())
	cmd.AddCommand(QueryPublicCapability())

	return cmd
}

func QueryEntitlementsOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner",
		Short: "Query the submodule's owner",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := entitlements.NewQueryClient(clientCtx)

			res, err := queryClient.Owner(context.Background(), &entitlements.QueryOwner{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func QueryPaused() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paused",
		Short: "Query if the module is paused",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := entitlements.NewQueryClient(clientCtx)

			res, err := queryClient.Paused(context.Background(), &entitlements.QueryPaused{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func QueryPublicCapabilities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "public-capabilities",
		Short: "Query for all public capabilities",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := entitlements.NewQueryClient(clientCtx)

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.PublicCapabilities(context.Background(), &entitlements.QueryPublicCapabilities{
				Pagination: pagination,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "public capabilities")

	return cmd
}

func QueryPublicCapability() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "public-capability [method]",
		Short: "Query if a method is public",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := entitlements.NewQueryClient(clientCtx)

			res, err := queryClient.PublicCapability(context.Background(), &entitlements.QueryPublicCapability{
				Method: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
