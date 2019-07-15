package parcel

import (
	"encoding/hex"
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
	data, err := storage.Inspect(args[0])
	if err != nil {
		fmt.Println("Error inspecting:", err)
		return nil
	}

	displayData := hex.EncodeToString(data)
	fmt.Println("Inspected data as a hex-encoded string:", displayData)

	return nil
}
