package query

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
)

var VCCmd = &cobra.Command{
	Use:   "vc <vcid>",
	Short: "Verifiable credential",
	Args:  cobra.MinimumNArgs(1),
	RunE:  vcFunc,
}

func vcFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryVC(args[0])
	if err != nil {
		return err
	}

	if rpc.DryRun {
		return nil
	}

	if asJson {
		fmt.Println(string(res))
		return nil
	}

	if res == nil || len(res) == 0 || string(res) == "null" {
		fmt.Println("not found")
		return nil
	}

	fmt.Printf("Verifiable credential: %s\n", string(res))

	return nil
}
