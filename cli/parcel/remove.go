package parcel

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/storage"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove <parcelID>",
	Short: "Remove data parcel with parcelID",
	Args:  cobra.MinimumNArgs(1),
	RunE:  removeFunc,
}

func removeFunc(cmd *cobra.Command, args []string) error {
	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	data, err := storage.Remove(args[0], key)
	if err != nil {
		fmt.Println("Error removing:", err)
		return nil
	}

	displayData := hex.EncodeToString(data)
	fmt.Println("Removed data as a hex-encoded string:", displayData)

	return nil
}
