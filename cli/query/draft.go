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
	var de types.DraftEx
	err = json.Unmarshal(res, &de)
	if err != nil {
		return err
	}

	cfg, err := json.MarshalIndent(de.Draft.Config, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("proposer: %s\n", de.Draft.Proposer)
	fmt.Printf("config: %s\n", cfg)
	fmt.Printf("desc: \"%s\"\n", de.Draft.Desc)
	fmt.Printf("open_count: %d\n", de.Draft.OpenCount)
	fmt.Printf("close_count: %d\n", de.Draft.CloseCount)
	fmt.Printf("apply_count: %d\n", de.Draft.ApplyCount)
	fmt.Printf("deposit: %s\n", de.Draft.Deposit.String())
	fmt.Printf("tally_quorum: %s\n", de.Draft.TallyQuorum.String())
	fmt.Printf("tally_approve: %s\n", de.Draft.TallyApprove.String())
	fmt.Printf("tally_reject: %s\n", de.Draft.TallyReject.String())
	fmt.Printf("votes: ")
	if len(de.Votes) == 0 {
		fmt.Printf("none\n")
	}
	for i, v := range de.Votes {
		fmt.Printf("  \n%d. voter: %s, approve: %t\n", i+1, v.Voter, v.Vote.Approve)
	}

	return nil
}
