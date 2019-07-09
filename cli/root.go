package cli

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/parcel"
	"github.com/amolabs/amo-client-go/cli/query"
	"github.com/amolabs/amo-client-go/cli/tx"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
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
		statusCmd,
		util.LineBreak,
		query.Cmd,
		tx.Cmd,
		parcel.Cmd,
		util.LineBreak,
	)
	RootCmd.PersistentFlags().StringP("rpc", "r", "0.0.0.0:26657", "ip:port")
	RootCmd.PersistentFlags().BoolP("json", "j", false, "output as json")
	RootCmd.PersistentFlags().StringP("user", "u", "", "username")
	RootCmd.PersistentFlags().StringP("pass", "p", "", "passphrase")
}

func readGlobalFlags(cmd *cobra.Command, args []string) {
	rpcArg, err := cmd.Flags().GetString("rpc")
	if err == nil {
		rpc.RpcRemote = "http://" + rpcArg
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
