package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var DraftCmd = &cobra.Command{
	Use:   "draft <draft_id>",
	Short: "Get draft status",
	Args:  cobra.MinimumNArgs(1),
	RunE:  draftFunc,
}

func draftFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryDraft(args[0])
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
		fmt.Println("no draft")
		return nil
	}
	var draft types.Draft
	err = json.Unmarshal(res, &draft)
	if err != nil {
		return err
	}

	cfg, err := json.Marshal(draft.Config)
	if err != nil {
		return err
	}

	fmt.Println("proposer:", draft.Proposer)
	fmt.Println("config:", cfg)
	fmt.Println("desc:", draft.Desc)
	fmt.Println("open_count:", draft.OpenCount)
	fmt.Println("close_count:", draft.CloseCount)
	fmt.Println("applye_count:", draft.ApplyCount)
	fmt.Println("deposit:", draft.Deposit)
	fmt.Println("tally_quorum:", draft.TallyQuorum)
	fmt.Println("tally_approve", draft.TallyApprove)
	fmt.Println("tally_reject:", draft.TallyReject)
	fmt.Println("votes:")
	for i, v := range draft.Votes {
		fmt.Printf("%d. voter: %s, approve: %b\n", i+1, v.Voter, v.Vote.Approve)
	}

	return nil
}
