package tx

import (
	"encoding/json"

	"github.com/amolabs/amo-client-go/lib/config"
	"github.com/amolabs/amo-client-go/lib/rpc"
)

func GetLastHeight(path string) (string, error) {
	lastHeight := ""
	cfg, err := config.GetConfig(path)
	if err != nil {
		return lastHeight, err
	}

	lastHeight = cfg.GetLastHeight()
	if lastHeight == "" {
		// get lastheight from outside
		rawMsg, err := rpc.NodeStatus()
		if err != nil {
			return lastHeight, err
		}

		jsonMsg, err := json.Marshal(rawMsg.SyncInfo)
		if err != nil {
			return lastHeight, err
		}

		data := make(map[string]interface{})
		err = json.Unmarshal(jsonMsg, &data)
		if err != nil {
			return lastHeight, err
		}

		lastHeight = data["latest_block_height"].(string)
		cfg.SetLastHeight(lastHeight)
		cfg.Save()
	}

	return lastHeight, nil
}

func SetLastHeight(path, lastHeight string) error {
	cfg, err := config.GetConfig(path)
	if err != nil {
		return err
	}

	cfg.SetLastHeight(lastHeight)

	err = cfg.Save()
	if err != nil {
		return err
	}

	return nil
}
