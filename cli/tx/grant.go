package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var GrantCmd = &cobra.Command{
	Use:   "grant <parcel_id> <address> <key_custody>",
	Short: "Grant a parcel permission",
	Args:  cobra.MinimumNArgs(3),
	RunE:  grantFunc,
}

func grantFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	extra, err := cmd.Flags().GetString("extra")
	if err != nil {
		return err
	}

	result, err := rpc.Grant(args[0], args[1], args[2], extra, key, Fee, Height)
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
	GrantCmd.PersistentFlags().String("extra", "null", "extra info")
}
