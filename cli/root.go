package cli

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
)

var LineBreak = &cobra.Command{Run: func(*cobra.Command, []string) {}}

var RootCmd = &cobra.Command{
	Use:              "amocli",
	Short:            "AMO blockchain console",
	PersistentPreRun: loadConfig,
}

func init() {
	cobra.EnableCommandSorting = false

	RootCmd.AddCommand(
		versionCmd,
		keyCmd,
		LineBreak,
		statusCmd,
		queryCmd,
		txCmd,
		parcelCmd,
		LineBreak,
	)
	RootCmd.PersistentFlags().StringP("rpc", "r", "0.0.0.0:26657",
		"node_ip:port")
	RootCmd.PersistentFlags().BoolP("json", "j", false,
		"output as json")
}

func loadConfig(cmd *cobra.Command, args []string) {
	rpcArg, err := cmd.Flags().GetString("rpc")
	if err == nil {
		rpc.RpcRemote = "http://" + rpcArg
	}
}
