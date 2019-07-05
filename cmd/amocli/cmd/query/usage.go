package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amoabci/amo/types"
)

var UsageCmd = &cobra.Command{
	Use:   "usage <buyer_address> <parcel_id>",
	Short: "Granted parcel usage",
	Args:  cobra.MinimumNArgs(2),
	RunE:  usageFunc,
}

func usageFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryUsage(args[0], args[1])
	if err != nil {
		return err
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("no usage")
	} else {
		var usage types.UsageValue
		err = json.Unmarshal(res, &usage)
		if err != nil {
			return err
		}
		fmt.Printf("custody: %s\nexpire: %s\n", usage.Custody, usage.Exp)
	}

	return nil
}
