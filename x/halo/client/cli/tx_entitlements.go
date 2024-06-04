package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
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

	return cmd
}
