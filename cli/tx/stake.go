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

	// This step is for user convenience. First, try base64 encoding, if it
	// succeeds encode it as hex again and send. If it fails assume hex
	// encoding and send it as is.
	var val string
	bin, err := base64.StdEncoding.DecodeString(args[0])
	if err == nil {
		val = hex.EncodeToString(bin)
	} else {
		val = args[0]
	}

	result, err := rpc.Stake(val, args[1], key)
	if err != nil {
		return err
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
