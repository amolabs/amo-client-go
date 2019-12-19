package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/config"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var StatusCmd = &cobra.Command{
	Use:   "node",
	Short: "Show status of AMO node",
	Long:  "Show status of AMO node including node info, pubkey, latest block hash, app hash, block height and time",
	RunE:  statusFunc,
}

func statusFunc(cmd *cobra.Command, args []string) error {
	rawMsg, err := rpc.NodeStatus()
	if err != nil {
		return err
	}

	if rpc.DryRun {
		return nil
	}

	jsonMsg, err := json.Marshal(rawMsg.SyncInfo)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(jsonMsg, &data)
	if err != nil {
		return err
	}

	lastHeight := data["latest_block_height"].(string)
	cfg, err := config.GetConfig(util.DefaultConfigFilePath())
	if err != nil {
		return err
	}

	if lastHeight == "0" {
		lastHeight = "1"
	}

	cfg.SetLastHeight(lastHeight)
	cfg.Save()

	resultJSON, err := json.MarshalIndent(rawMsg, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(resultJSON))

	return nil

}

func init() {
	// init here if needed
}
