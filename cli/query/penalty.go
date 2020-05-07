package query

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var PenaltyCmd = &cobra.Command{
	Use:   "penalty <block_height | address>",
	Short: "Penalty history of given data",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  penaltyFunc,
}

func penaltyFunc(cmd *cobra.Command, args []string) error {
	var (
		res             []byte
		err             error
		height, address string
		penalties       []types.PenaltyInfo
	)

	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	switch len(args) {
	case 1:
		if _, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			res, err = rpc.QueryBlockPenalty(args[0])
		} else {
			res, err = rpc.QueryAddressPenalty(args[0])
		}
	case 2:
		if _, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			height = args[0]
			address = args[1]
		} else if _, err := strconv.ParseInt(args[1], 10, 64); err == nil {
			height = args[1]
			address = args[0]
		} else {
			return fmt.Errorf("unavaialbe arguments")
		}

		res, err = rpc.QueryPenalty(height, address)
	}

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
		fmt.Println("no penalty")
		return nil
	}

	err = json.Unmarshal(res, &penalties)
	if err != nil {
		return err
	}

	sort.Slice(penalties, func(i, j int) bool {
		return penalties[i].BlockHeight < penalties[j].BlockHeight
	})

	for _, penalty := range penalties {
		fmt.Printf("block height: %d, address: %s, amount: %s\n",
			penalty.BlockHeight, penalty.Address, penalty.Amount.String())
	}

	return nil
}
