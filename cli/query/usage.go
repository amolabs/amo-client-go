package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
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

	if rpc.DryRun {
		return nil
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("no usage")
		return nil
	}

	var usage types.UsageEx
	err = json.Unmarshal(res, &usage)
	if err != nil {
		return err
	}

	fmt.Printf("custody: %s\n", usage.Custody)
	fmt.Printf("buyer: %s\n", usage.Buyer)
	fmt.Printf("extra: %s\n", usage.Extra)

	return nil
}
