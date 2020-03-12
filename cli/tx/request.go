package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var RequestCmd = &cobra.Command{
	Use:   "request <parcel_id> <amount>",
	Short: "Request a parcel permission with payment",
	Args:  cobra.MinimumNArgs(2),
	RunE:  requestFunc,
}

func requestFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	dealer, err := cmd.Flags().GetString("dealer")
	if err != nil {
		return err
	}

	dealerFee, err := cmd.Flags().GetString("dealer_fee")
	if err != nil {
		return err
	}

	extra, err := cmd.Flags().GetString("extra")
	if err != nil {
		return err
	}

	result, err := rpc.Request(args[0], args[1], dealer, dealerFee, extra, key, Fee, Height)
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

func init() {
	RequestCmd.PersistentFlags().String("dealer", "", "dealer address")
	RequestCmd.PersistentFlags().String("dealer_fee", "", "fee to pay for dealer")
	RequestCmd.PersistentFlags().String("extra", "", "extra info")
}
