package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var RequestCmd = &cobra.Command{
	Use:   "request <buyer_address> <parcel_id>",
	Short: "Requested parcel usage",
	Args:  cobra.MinimumNArgs(2),
	RunE:  requestFunc,
}

func requestFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryRequest(args[0], args[1])
	if err != nil {
		return err
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("no request")
	} else {
		var request types.Request
		err = json.Unmarshal(res, &request)
		if err != nil {
			return err
		}
		fmt.Printf("payment: %s\nexpire: %s\n", request.Payment, request.Exp)
	}

	return nil
}
