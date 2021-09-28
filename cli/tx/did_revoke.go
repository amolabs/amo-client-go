package tx

import (
	"encoding/json"
	"fmt"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/spf13/cobra"
)

var DIDRevokeCmd = &cobra.Command{
	Use:   "did.revoke <vcid>",
	Short: "Revoke a verifiable credential",
	Args:  cobra.MinimumNArgs(1),
	RunE:  didDismissFunc,
}

func didRevokeFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.DIDRevoke(args[0], key, Fee, Height)
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

	// TODO: rich output

	return nil
}
