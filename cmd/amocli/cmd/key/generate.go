package key

import (
	"fmt"

	"github.com/spf13/cobra"

	cliutil "github.com/amolabs/amo-client-go/cmd/amocli/util"
	"github.com/amolabs/amo-client-go/lib/keys"
	"github.com/amolabs/amo-client-go/lib/util"
)

var GenCmd = &cobra.Command{
	Use:   "generate <username>",
	Short: "Generate a key with a specified username",
	Args:  cobra.MinimumNArgs(1),
	RunE:  genFunc,
}

func init() {
	cmd := GenCmd
	cmd.Flags().SortFlags = false
	cmd.Flags().BoolP("encrypt", "e", true, "encrypt the private key with passphrase")
	cmd.Flags().StringP("seed", "s", "", "optional seed string")
}

func genFunc(cmd *cobra.Command, args []string) error {
	username := args[0]
	keyFile := util.DefaultKeyFilePath()
	flags := cmd.Flags()

	encrypt, err := flags.GetBool("encrypt")
	if err != nil {
		return err
	}

	seed, err := flags.GetString("seed")
	if err != nil {
		return err
	}

	var passphrase []byte

	if encrypt {
		b, err := cliutil.PromptPassphrase()
		if err != nil {
			return err
		}
		passphrase = []byte(b)
	}

	kr, err := keys.GetKeyRing(keyFile)
	if err != nil {
		return err
	}
	_, err = kr.GenerateNewKey(username, passphrase, encrypt, seed)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully generated the key with username: %s\n", username)

	return nil
}
