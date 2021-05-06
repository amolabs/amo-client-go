package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var TransferCmd = &cobra.Command{
	Use:   "transfer <address> <amount>|<parcel_id>",
	Short: "Transfer the specified amount of money or parcel to <address>",
	Args:  cobra.MinimumNArgs(2),
	//	PersistentPreRun: readGlobalFlags,
	RunE: transferFunc,
}

func transferFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	udc, err := cmd.Flags().GetUint32("udc")
	if err != nil {
		return err
	}

	isParcel, err := cmd.Flags().GetBool("parcel")
	if err != nil {
		return err
	}

	result, err := rpc.Transfer(udc, args[0], args[1], isParcel, key, Fee, Height)
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
	TransferCmd.PersistentFlags().Uint32("udc", uint32(0), "specify udc id if necessary")
	TransferCmd.PersistentFlags().Bool("parcel", false, "when true, transfer parcel instead of coin")
}
