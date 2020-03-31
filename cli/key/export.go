package key

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/keys"
)

var ExportCmd = &cobra.Command{
	Use:   "export <username>",
	Short: "Import a raw private key from base64-formatted string",
	Args:  cobra.MinimumNArgs(1),
	RunE:  exportFunc,
}

func init() {
}

func exportFunc(cmd *cobra.Command, args []string) error {
	username := args[0]
	keyFile := util.DefaultKeyFilePath()

	kr, err := keys.GetKeyRing(keyFile)
	if err != nil {
		return err
	}
	key := kr.GetKey(username)
	if key == nil {
		return errors.New("Failed to get key")
	}
	if key.Encrypted {
		b, err := util.PromptPassphrase()
		if err != nil {
			return err
		}
		passphrase := []byte(b)
		err = key.Decrypt(passphrase)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Private key for user '%s'\n  Hex: %s\n  Base64: %s\n",
		username,
		hex.EncodeToString(key.PrivKey),
		base64.StdEncoding.EncodeToString(key.PrivKey),
	)

	return nil
}
