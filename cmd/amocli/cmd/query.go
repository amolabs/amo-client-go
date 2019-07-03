package cmd

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amoabci/cmd/amocli/cmd/query"
)

/* Commands (expected hierarchy)
 *
 * amocli |- query |- balance <address>
 *                 |
 *                 |- stake <address>
 *                 |- delegate --holder <address> --delegator <address>
 *                 |
 *                 |- parcel <parcelID>
 *                 |- request --buyer <address> --target <parcelID>
 *                 |- usage --buyer <address> --target <parcelID>
 */

var queryCmd = &cobra.Command{
	Use:     "query",
	Aliases: []string{"q"},
	Short:   "Query AMO blockchain data",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	queryCmd.AddCommand(
		query.BalanceCmd,
		query.StakeCmd,
		query.DelegateCmd,
		LineBreak,
		query.ParcelCmd,
		query.RequestCmd,
		query.UsageCmd,
	)
}
