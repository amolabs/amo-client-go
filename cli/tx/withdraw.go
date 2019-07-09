package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var WithdrawCmd = &cobra.Command{
	Use:   "withdraw <amount>",
	Short: "Withdraw all or part of the AMO coin locked as a stake",
	Args:  cobra.MinimumNArgs(1),
	RunE:  withdrawFunc,
}

func withdrawFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Withdraw(args[0], key)
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
