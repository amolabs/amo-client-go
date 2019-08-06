package parcel

import (
	"encoding/hex"
	"encoding/json"
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
	var err error

	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	var data []byte
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

	res, err := storage.Upload(data, key)
	if err != nil {
		if asJson {
			fmt.Println(err)
		} else {
			fmt.Println("Error uploading:", err)
		}
		return nil
	}

	if asJson {
		fmt.Println(string(res))
	} else {
		var parcel struct {
			id string `json:"id"`
		}
		err = json.Unmarshal([]byte(res), &parcel)
		if err != nil {
			return err
		}
		fmt.Println("New parcel ID:", parcel.id)
	}

	return nil
}
