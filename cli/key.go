package cli

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
)

var keyCmd = &cobra.Command{
	Use:     "key",
	Aliases: []string{"k"},
	Short:   "Manage local keyring",
}

func init() {
	cmd := keyCmd

	cmd.AddCommand(
		key.ListCmd,
		key.ImportCmd,
		key.GenCmd,
		key.RemoveCmd,
	)
}
