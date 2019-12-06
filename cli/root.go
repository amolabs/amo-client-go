package cli

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/parcel"
	"github.com/amolabs/amo-client-go/cli/query"
	"github.com/amolabs/amo-client-go/cli/tx"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/storage"
)

var RootCmd = &cobra.Command{
	Use:              "amocli",
	Short:            "AMO blockchain console",
	PersistentPreRun: readGlobalFlags,
}

func init() {
	cobra.EnableCommandSorting = false

	RootCmd.AddCommand(
		versionCmd,
		key.Cmd,
		util.LineBreak,
		query.Cmd,
		tx.Cmd,
		parcel.Cmd,
		util.LineBreak,
	)
	RootCmd.PersistentFlags().StringP("rpc", "r", "0.0.0.0:26657", "ip:port")
	// TODO: change shorcut or reorganize global flags
	RootCmd.PersistentFlags().String("sto", "0.0.0.0:80", "ip:port")
	RootCmd.PersistentFlags().BoolP("json", "j", false, "output as json")
	RootCmd.PersistentFlags().BoolP("dry", "d", false, "dry run")
	RootCmd.PersistentFlags().StringP("user", "u", "", "username")
	RootCmd.PersistentFlags().StringP("pass", "p", "", "passphrase")
}

func readGlobalFlags(cmd *cobra.Command, args []string) {
	rpcArg, err := cmd.Flags().GetString("rpc")
	if err == nil {
		rpc.RpcRemote = "http://" + rpcArg
	}
	stoArg, err := cmd.Flags().GetString("sto")
	if err == nil {
		storage.Endpoint = "http://" + stoArg
	}
	dryRun, err := cmd.Flags().GetBool("dry")
	if err == nil {
		rpc.DryRun = dryRun
	}
	feeArg, err := cmd.Flags().GetString("fee")
	if err == nil {
		tx.Fee = feeArg
	}
	broadcastOptionArg, err := cmd.Flags().GetString("broadcast")
	if err == nil {
		rpc.TxBroadcastOption = broadcastOptionArg
	}
	username, err := cmd.Flags().GetString("user")
	if err == nil {
		key.Username = username
	}
	passphrase, err := cmd.Flags().GetString("pass")
	if err == nil {
		key.Passphrase = passphrase
	}
}
