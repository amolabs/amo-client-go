package tx

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
)

var Cmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"t"},
	Short:   "Send signed transactions",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(
		TransferCmd,
		util.LineBreak,
		StakeCmd,
		WithdrawCmd,
		DelegateCmd,
		RetractCmd,
		util.LineBreak,
		RegisterCmd,
		RequestCmd,
		GrantCmd,
		util.LineBreak,
		DiscardCmd,
		CancelCmd,
		RevokeCmd,
	)
}
