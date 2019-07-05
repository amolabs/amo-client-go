package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amoabci/amo/types"
)

var ParcelCmd = &cobra.Command{
	Use:   "parcel <parcelID>",
	Short: "Data parcel detail",
	Args:  cobra.MinimumNArgs(1),
	RunE:  parcelFunc,
}

func parcelFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryParcel(args[0])
	if err != nil {
		return err
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("no parcel")
	} else {
		var parcel types.ParcelValue
		err = json.Unmarshal(res, &parcel)
		if err != nil {
			return err
		}
		fmt.Printf("owner: %s\ncustody: %s\n", parcel.Owner, parcel.Custody)
	}

	return nil
}
