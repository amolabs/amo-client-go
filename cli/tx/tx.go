package tx

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
)

var (
	Fee    string
	Height string
)

var Cmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"t"},
	Short:   "Send signed transactions",
}

func init() {
	Cmd.AddCommand(
		TransferCmd,
		util.LineBreak,
		IssueCmd,
		BurnCmd,
		LockCmd,
		util.LineBreak,
		StakeCmd,
		WithdrawCmd,
		DelegateCmd,
		RetractCmd,
		util.LineBreak,
		ProposeCmd,
		VoteCmd,
		util.LineBreak,
		SetupCmd,
		CloseCmd,
		util.LineBreak,
		RegisterCmd,
		RequestCmd,
		GrantCmd,
		util.LineBreak,
		DiscardCmd,
		CancelCmd,
		RevokeCmd,
		util.LineBreak,
		DIDClaimCmd,
		DIDDismissCmd,
		DIDIssueCmd,
		DIDRevokeCmd,
	)

	Cmd.PersistentFlags().StringP("fee", "f", "0", "fee for tx")
	Cmd.PersistentFlags().String("height", "", "height for block binding tx")
	Cmd.PersistentFlags().StringP("broadcast", "b", "sync", "options(commit, sync, async) for broadcast method")
}
