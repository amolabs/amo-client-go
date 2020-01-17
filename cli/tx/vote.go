package tx

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var VoteCmd = &cobra.Command{
	Use:   "vote <draft_id> <approve>",
	Short: "Vote on a draft",
	Args:  cobra.MinimumNArgs(2),
	RunE:  voteFunc,
}

func voteFunc(cmd *cobra.Command, args []string) error {
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

	apBool, err := strconv.ParseBool(args[1])
	if err != nil {
		return err
	}

	result, err := rpc.Vote(args[0], apBool, key, Fee, lastHeight)
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

	return nil
}
