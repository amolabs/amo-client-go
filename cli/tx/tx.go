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
	Cmd.PersistentPreRun = preRunChain
	Cmd.PersistentFlags().StringP("user", "u", "", "username")
	Cmd.PersistentFlags().StringP("pass", "p", "",
		"passphrase of an encrypted key")
}

func preRunChain(cmd *cobra.Command, args []string) {
	// If this function runs, it means that no children commands designated
	// persistentPreRun. In that case, scan upward and search first occurrence
	// of persistentPreRun and run it first. This is necessary because cobra
	// just scans the first occurrence of persistentPreRun from the leaf
	// command, but we need chain of persistentPreRun.
	beep := false
	for c := cmd; c != nil; c = c.Parent() {
		if c == Cmd {
			beep = true
			continue
		}
		if run := c.PersistentPreRun; beep && run != nil {
			run(cmd, args)
			break
		}
	}

	readUserPass(cmd, args)
}

func readUserPass(cmd *cobra.Command, args []string) {
	username, err := cmd.Flags().GetString("user")
	if err == nil {
		Username = username
	}

	passphrase, err := cmd.Flags().GetString("pass")
	if err == nil {
		Passphrase = passphrase
	}
}
