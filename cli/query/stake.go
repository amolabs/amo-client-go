package query

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var StakeCmd = &cobra.Command{
	Use:   "stake <address>",
	Short: "Stake of an account",
	Args:  cobra.MinimumNArgs(1),
	RunE:  stakeFunc,
}

func stakeFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	res, err := rpc.QueryStake(args[0])
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
		fmt.Println("no stake")
	} else {
		var stake types.Stake
		err = json.Unmarshal(res, &stake)
		if err != nil {
			return err
		}
		fmt.Printf("amount: %s\nvalidator pubkey: %s\n",
			stake.Amount.String(), hex.EncodeToString(stake.Validator))
		for i, d := range stake.Delegates {
			fmt.Printf("  delegate %2d: %s from %s\n",
				i+1, d.Amount.String(), d.Delegator)
		}
	}

	return nil
}
