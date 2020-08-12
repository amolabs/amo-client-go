package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var RequestCmd = &cobra.Command{
	Use:   "request <parcel_id> <recipient>",
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

	if rpc.DryRun {
		return nil
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("no request")
		return nil
	}

	var request types.RequestEx
	err = json.Unmarshal(res, &request)
	if err != nil {
		return err
	}

	fmt.Printf("payment: %s\n", request.Payment.String())
	fmt.Printf("agency: %s\n", request.Agency)
	fmt.Printf("dealer: %s\n", request.Dealer)
	fmt.Printf("dealer_fee: %s\n", request.DealerFee.String())
	fmt.Printf("extra: %s\n", request.Extra)
	fmt.Printf("recipient: %s\n", request.Recipient)

	return nil
}
