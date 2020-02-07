package tx

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var IssueCmd = &cobra.Command{
	Use:   "issue <udc_id> <amount>",
	Short: "Issue a UDC with its id and total issue amount",
	Args:  cobra.MinimumNArgs(2),
	RunE:  issueFunc,
}

func issueFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	lastHeight, err := GetLastHeight(util.DefaultConfigFilePath())
	if err != nil {
		return err
	}

	desc, err := cmd.Flags().GetString("desc")
	if err != nil {
		return err
	}

	operators, err := cmd.Flags().GetStringSlice("operators")
	if err != nil {
		return err
	}

	result, err := rpc.Issue(args[0], args[1], desc, operators, key, Fee, lastHeight)
	if err != nil {
		return err
	}

	if result.Height != "0" {
		SetLastHeight(util.DefaultConfigFilePath(), result.Height)
	}

	if asJson {
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return err
		}

		fmt.Println(string(resultJSON))
	}

	// TODO: rich output
	return nil
}

func init() {
	IssueCmd.PersistentFlags().String("desc", "", "description")
	IssueCmd.PersistentFlags().StringSliceP("operators", "o", []string{}, "operator addresses")
}
