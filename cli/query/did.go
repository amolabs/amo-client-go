package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var DIDCmd = &cobra.Command{
	Use:   "did <did>",
	Short: "DID document",
	Args:  cobra.MinimumNArgs(1),
	RunE:  didFunc,
}

func didFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryDID(args[0])
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

	var entry types.DIDEntry
	err = json.Unmarshal(res, &entry)
	if err != nil {
		return err
	}

	fmt.Printf("DID document: %s\n", string(entry.Document))
	fmt.Printf("Metadata: %s\n", string(entry.Meta))

	return nil
}
