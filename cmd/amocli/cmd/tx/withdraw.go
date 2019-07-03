package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/util"
	atypes "github.com/amolabs/amoabci/amo/types"
)

var WithdrawCmd = &cobra.Command{
	Use:   "withdraw <amount>",
	Short: "Withdraw all or part of the AMO coin locked as a stake",
	Args:  cobra.MinimumNArgs(1),
	RunE:  withdrawFunc,
}

func withdrawFunc(cmd *cobra.Command, args []string) error {
	amount, err := new(atypes.Currency).SetString(args[0], 10)
	if err != nil {
		return err
	}

	key, err := GetRawKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	result, err := rpc.Withdraw(amount, key)
	if err != nil {
		return err
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	fmt.Println(string(resultJSON))

	return nil
}
