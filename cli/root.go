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
	PersistentPreRun: loadConfig,
}

func init() {
	cobra.EnableCommandSorting = false

	RootCmd.AddCommand(
		versionCmd,
		key.Cmd,
		util.LineBreak,
		statusCmd,
		query.Cmd,
		tx.Cmd,
		parcel.Cmd,
		util.LineBreak,
	)
	RootCmd.PersistentFlags().StringP("rpc", "r", "0.0.0.0:26657", "ip:port")
	RootCmd.PersistentFlags().BoolP("json", "j", false, "output as json")
}

func loadConfig(cmd *cobra.Command, args []string) {
	rpcArg, err := cmd.Flags().GetString("rpc")
	if err == nil {
		rpc.RpcRemote = "http://" + rpcArg
	}
}
