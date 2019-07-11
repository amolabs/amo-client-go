package parcel

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/storage"
)

var DownloadCmd = &cobra.Command{
	Use:   "download <parcelID>",
	Short: "Download data parcel with parcelID",
	Args:  cobra.MinimumNArgs(1),
	RunE:  downloadFunc,
}

func downloadFunc(cmd *cobra.Command, args []string) error {
	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	data, err := storage.Download(args[0], key)
	if err != nil {
		fmt.Println("Error downloading:", err)
		return nil
	}

	displayData := hex.EncodeToString(data)
	fmt.Println("Downloaded data as a hex-encoded string:", displayData)

	return nil
}
