package parcel

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

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

func init() {
	cmd := DownloadCmd
	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("file", "f", "", "file to save")
}

func downloadFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	doSave := false
	filename, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	if len(filename) > 0 {
		doSave = true
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	data, err := storage.Download(args[0], key)
	if err != nil {
		if asJson {
			fmt.Println(err)
		} else {
			fmt.Println("Error downloading:", err)
		}
		return nil
	}

	if doSave {
		err = ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Downloaded data has been saved to the file:", filename)
	} else {
		displayData := hex.EncodeToString(data)
		fmt.Println("Downloaded data as a hex-encoded string:", displayData)
	}

	return nil
}
