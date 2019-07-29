package key

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "key",
	Aliases: []string{"k"},
	Short:   "Manage local keyring",
}

func init() {
	Cmd.AddCommand(
		ListCmd,
		ImportCmd,
		ExportCmd,
		GenCmd,
		RemoveCmd,
	)
}
