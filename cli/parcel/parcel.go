package parcel

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "parcel",
	Short: "Data parcel operations",
}

func init() {
	Cmd.AddCommand(
		UploadCmd,
		RetrieveCmd,
		QueryCmd,
	)
}
