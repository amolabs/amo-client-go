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
	Short: "Draft status",
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
	var draft types.DraftEx
	err = json.Unmarshal(res, &draft)
	if err != nil {
		return err
	}

	cfg, err := json.MarshalIndent(draft.Config, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("proposer: %s\n", draft.Proposer)
	fmt.Printf("config: %s\n", cfg)
	fmt.Printf("desc: \"%s\"\n", draft.Desc)
	fmt.Printf("open_count: %d\n", draft.OpenCount)
	fmt.Printf("close_count: %d\n", draft.CloseCount)
	fmt.Printf("apply_count: %d\n", draft.ApplyCount)
	fmt.Printf("deposit: %s\n", draft.Deposit.String())
	fmt.Printf("tally_quorum: %s\n", draft.TallyQuorum.String())
	fmt.Printf("tally_approve: %s\n", draft.TallyApprove.String())
	fmt.Printf("tally_reject: %s\n", draft.TallyReject.String())
	fmt.Printf("votes: ")
	if len(draft.Votes) == 0 {
		fmt.Printf("none\n")
	} else {
		fmt.Print("\n")
	}
	for i, v := range draft.Votes {
		fmt.Printf("  %d. voter: %s, approve: %t\n", i+1, v.Voter, v.Vote.Approve)
	}

	return nil
}
