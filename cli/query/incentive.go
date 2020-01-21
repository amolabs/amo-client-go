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

var IncentiveCmd = &cobra.Command{
	Use:   "incentive <block_height | address>",
	Short: "Incentive history of given data",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  incentiveFunc,
}

func incentiveFunc(cmd *cobra.Command, args []string) error {
	var (
		res             []byte
		err             error
		height, address string
		incentives      []types.IncentiveInfo
	)

	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	switch len(args) {
	case 1:
		if _, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			res, err = rpc.QueryBlockIncentive(args[0])
		} else {
			res, err = rpc.QueryAddressIncentive(args[0])
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

		res, err = rpc.QueryIncentive(height, address)
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
		fmt.Println("no incentive")
		return nil
	}

	err = json.Unmarshal(res, &incentives)
	if err != nil {
		return err
	}

	sort.Slice(incentives, func(i, j int) bool {
		return incentives[i].BlockHeight < incentives[j].BlockHeight
	})

	for _, incentive := range incentives {
		fmt.Printf("block height: %d, address: %s, amount: %s\n",
			incentive.BlockHeight, incentive.Address, incentive.Amount.String())
	}

	return nil
}
