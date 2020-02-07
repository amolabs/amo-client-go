package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var RegisterCmd = &cobra.Command{
	Use:   "register <parcel_id> <key_custody>",
	Short: "Register a parcel with extra information",
	Args:  cobra.MinimumNArgs(2),
	RunE:  registerFunc,
}

func registerFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	proxy, err := cmd.Flags().GetString("proxy")
	if err != nil {
		return err
	}

	extra, err := cmd.Flags().GetString("extra")
	if err != nil {
		return err
	}

	result, err := rpc.Register(args[0], args[1], proxy, extra, key, Fee, Height)
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
	RegisterCmd.PersistentFlags().String("proxy", "", "proxy account of parcel")
	RegisterCmd.PersistentFlags().String("extra", "null", "extra info")
}
