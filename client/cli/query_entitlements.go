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
	"github.com/noble-assets/halo/v3/types/entitlements"
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
	cmd.AddCommand(QueryPublicCapability())
	cmd.AddCommand(QueryRoleCapability())

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

func QueryRoleCapability() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "role-capability [method]",
		Short: "Query roles for a specific method",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := entitlements.NewQueryClient(clientCtx)

			res, err := queryClient.RoleCapability(context.Background(), &entitlements.QueryRoleCapability{
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

func QueryUserCapability() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-capability [address]",
		Short: "Query roles for a specific address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := entitlements.NewQueryClient(clientCtx)

			res, err := queryClient.UserCapability(context.Background(), &entitlements.QueryUserCapability{
				Address: args[0],
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
