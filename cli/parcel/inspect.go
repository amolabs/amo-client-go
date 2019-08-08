package parcel

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/storage"
)

var InspectCmd = &cobra.Command{
	Use:   "inspect <parcelID>",
	Short: "Inspect data parcel with parcelID",
	Args:  cobra.MinimumNArgs(1),
	RunE:  inspectFunc,
}

func inspectFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	resJson, err := storage.Inspect(args[0])
	if err != nil {
		if asJson {
			fmt.Println(err)
		} else {
			fmt.Println("Error inspecting:", err)
		}
		return nil
	}

	if asJson {
		fmt.Println(string(resJson))
	} else {
		var res struct {
			Metadata json.RawMessage `json:"metadata"`
		}
		err = json.Unmarshal(resJson, &res)
		if err != nil {
			return err
		}
		fmt.Println("Metadata:", string(res.Metadata))
	}

	return nil
}
