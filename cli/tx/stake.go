package tx

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/cli/key"
	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

var StakeCmd = &cobra.Command{
	Use:   "stake <validator_pubkey> <amount>",
	Short: "Lock AMO coin and acquire a stake with a validator key",
	Args:  cobra.MinimumNArgs(2),
	RunE:  stakeFunc,
}

func stakeFunc(cmd *cobra.Command, args []string) error {
	asJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		return err
	}

	key, err := key.GetUserKey(util.DefaultKeyFilePath())
	if err != nil {
		return err
	}

	// This step is for user convenience. First, try decoding it in hex, if it
	// succeeds send it as it is. If it fails assume base64 encoding and decode
	// it in base64 and encode it in hex and send it.
	var val string
	_, err = hex.DecodeString(args[0])
	if err == nil {
		val = args[0]
	} else {
		bin, err := base64.StdEncoding.DecodeString(args[0])
		if err != nil {
			return err
		}
		val = hex.EncodeToString(bin)
	}

	result, err := rpc.Stake(val, args[1], key, Fee, Height)
	if err != nil {
		return err
	}

	if rpc.DryRun {
		return nil
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
