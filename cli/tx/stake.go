package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var StakeCmd = &cobra.Command{
	Use:   "stake <validator_pubkey> <amount>",
	Short: "Lock AMO coin and acquire a stake with a validator key",
	Args:  cobra.MinimumNArgs(2),
	RunE:  stakeFunc,
}

func stakeFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Stake(args[0], args[1], key)
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