package key

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/keys"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all of keys stored on the local storage",
	Args:  cobra.NoArgs,
	RunE:  listFunc,
}

func init() {
	cmd := ListCmd
	cmd.Flags().BoolP("pubkey", "k", false, "show public key")
}

func listFunc(cmd *cobra.Command, args []string) error {
	keyFile := util.DefaultKeyFilePath()
	flags := cmd.Flags()

	kr, err := keys.GetKeyRing(keyFile)
	if err != nil {
		return err
	}

	withPubKey, err := flags.GetBool("pubkey")

	kr.PrintKeyList(withPubKey)

	return nil
}
