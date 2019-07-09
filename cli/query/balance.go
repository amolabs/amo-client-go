package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amoabci/amo/types"
)

var BalanceCmd = &cobra.Command{
	Use:   "balance <address>",
	Short: "Coin balance of an account",
	Args:  cobra.MinimumNArgs(1),
	RunE:  balanceFunc,
}

func balanceFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	// TODO: do some sanity check on client side
	res, err := rpc.QueryBalance(args[0])
	if err != nil {
		return err
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	var balance types.Currency
	err = json.Unmarshal([]byte(res), &balance)
	if err != nil {
		return err
	}
	fmt.Println(balance, "mote") // TODO: print AMO unit also

	return nil
}
