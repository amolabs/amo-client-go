package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var RequestCmd = &cobra.Command{
	Use:   "request <parcel_id> <amount>",
	Short: "Request a parcel permission with payment",
	Args:  cobra.MinimumNArgs(2),
	RunE:  requestFunc,
}

func requestFunc(cmd *cobra.Command, args []string) error {
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

	extra, err := cmd.Flags().GetString("extra")
	if err != nil {
		return err
	}

	result, err := rpc.Request(args[0], args[1], extra, key, Fee, lastHeight)
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

func init() {
	RequestCmd.PersistentFlags().String("extra", "null", "extra info")
}
