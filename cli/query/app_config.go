package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

var AppConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Show config of AMO abci",
	RunE:  appConfigFunc,
}

func appConfigFunc(cmd *cobra.Command, args []string) error {
	// TODO: do some sanity check on client side
	res, err := rpc.QueryAppConfig()
	if err != nil {
		return err
	}

	if rpc.DryRun {
		return nil
	}

	// TODO: cache app config somewhere in a storage
	var appConfig types.AMOAppConfig
	err = json.Unmarshal([]byte(res), &appConfig)
	if err != nil {
		return err
	}

	fmt.Println(string(res))

	return nil
}
