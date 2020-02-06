package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var BurnCmd = &cobra.Command{
	Use:   "burn <udc_id> <amount>",
	Short: "Burn certain amount of UDC",
	Args:  cobra.MinimumNArgs(2),
	RunE:  burnFunc,
}

func burnFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	lastHeight, err := GetLastHeight(util.DefaultConfigFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Burn(args[0], args[1], key, Fee, lastHeight)
	if err != nil {
		return err
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

	// TODO: rich output
	return nil
}
