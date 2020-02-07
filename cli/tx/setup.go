package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var SetupCmd = &cobra.Command{
	Use:   "setup <storage_id> <url> <registration_fee> <hosting_fee>",
	Short: "Setup a storage with url and paying fees",
	Args:  cobra.MinimumNArgs(4),
	RunE:  setupFunc,
}

func setupFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Setup(args[0], args[1], args[2], args[3], key, Fee, Height)
	if err != nil {
		return err
	}

	if rpc.DryRun {
		return nil
	}

	if result.Height != "0" {
		SetLastHeight(util.DefaultConfigFilePath(), result.Height)
	}

	if asJson {
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return err
		}

		fmt.Println(string(resultJSON))
	}

	return nil
}
