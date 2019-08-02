package parcel

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/storage"
)

var UploadCmd = &cobra.Command{
	Use:   "upload ( <hex> | --file <filename> )",
	Short: "Upload data parcel",
	Args:  cobra.NoArgs,
	RunE:  uploadFunc,
}

func init() {
	cmd := UploadCmd
	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("file", "f", "", "file to upload")
}

func uploadFunc(cmd *cobra.Command, args []string) error {
	var data []byte
	var err error
	if len(args) > 0 {
		data, err = hex.DecodeString(args[0])
		if err != nil {
			return err
		}
	} else {
		filename, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}
		if len(filename) == 0 {
			return errors.New("Either hex data or filename must be given.")
		}
		data, err = ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
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
