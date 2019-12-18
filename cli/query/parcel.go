package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
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

	if rpc.DryRun {
		return nil
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("no parcel")
	} else {
		var parcel types.Parcel
		err = json.Unmarshal(res, &parcel)
		if err != nil {
			return err
		}
		fmt.Printf("owner: %s\ncustody: %s\n", parcel.Owner, parcel.Custody)
		for i, r := range parcel.Requests {
			fmt.Printf("  requests %2d. payment: %s, expire: %s, buyer: %s\n",
				i+1, r.Payment.String(), r.Exp, r.Buyer)
		}
		for i, u := range parcel.Usages {
			fmt.Printf("  usages %2d. custody: %s, expire: %s, buyer: %s\n",
				i+1, u.Custody, u.Exp, u.Buyer)
		}
	}

	return nil
}
