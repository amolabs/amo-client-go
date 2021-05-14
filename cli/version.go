package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION represents the general version of this app
const VERSION = "v1.8.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	// init here if needed
}
