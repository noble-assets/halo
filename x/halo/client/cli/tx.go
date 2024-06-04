package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/noble-assets/halo/x/halo/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transactions commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetAggregatorTxCmd())
	cmd.AddCommand(GetEntitlementsTxCmd())

	// TODO: Add missing commands.

	return cmd
}
