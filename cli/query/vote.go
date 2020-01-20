package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var VoteCmd = &cobra.Command{
	Use:   "vote <draft_id> <address>",
	Short: "Get vote status of given voter address",
	Args:  cobra.MinimumNArgs(1),
	RunE:  voteFunc,
}

func voteFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryVote(args[0], args[1])
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
		fmt.Println("no vote")
		return nil
	}

	var vote types.Vote
	err = json.Unmarshal(res, &vote)
	if err != nil {
		return err
	}

	fmt.Println("draft_id:", args[0])
	fmt.Println("voter:", args[1])
	fmt.Println("approve:", vote.Approve)

	return nil
}
