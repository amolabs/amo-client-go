package parcel

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/storage"
)

var UploadCmd = &cobra.Command{
	Use:   "upload <hex>",
	Short: "Upload data parcel",
	Args:  cobra.MinimumNArgs(1),
	RunE:  uploadFunc,
}

func init() {
	//cmd := UploadCmd
	//cmd.Flags().SortFlags = false
	//cmd.Flags().StringP("qualifier", "q", "", "extra data info")
}

func uploadFunc(cmd *cobra.Command, args []string) error {
	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	data, err := hex.DecodeString(args[0])
	if err != nil {
		return err
	}
	parcelID, err := storage.Upload(data, key)
	if err != nil {
		return err
	}

	fmt.Println("New parcel ID:", parcelID)

	return nil
}
