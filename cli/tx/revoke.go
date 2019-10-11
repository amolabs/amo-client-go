package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var RevokeCmd = &cobra.Command{
	Use:   "revoke <parcel_id> <address>",
	Short: "Delete the usage of parcel in store/usage",
	Args:  cobra.MinimumNArgs(2),
	RunE:  revokeFunc,
}

func revokeFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Revoke(args[0], args[1], key, Fee)
	if err != nil {
		return err
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
