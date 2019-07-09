package cli

import (
	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/parcel"
)

var parcelCmd = &cobra.Command{
	Use:   "parcel",
	Short: "Data parcel operations",
}

func init() {
	parcelCmd.AddCommand(
		parcel.UploadCmd,
		parcel.RetrieveCmd,
		parcel.QueryCmd,
	)
}
