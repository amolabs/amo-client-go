package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/amolabs/amo-client-go/cli/util"
	"github.com/amolabs/amo-client-go/lib/rpc"
	"github.com/amolabs/amo-client-go/lib/types"
)

type ConfigSet struct {
	LastHeight string             `json:"last_height"`
	ABCIConfig types.AMOAppConfig `json:"ABCI_config"`
	// TODO: CLIConfig would be needed here in the future
}

type Config struct {
	filePath  string
	configSet ConfigSet
}

func GetConfig(path string) (*Config, error) {
	cfg := new(Config)
	cfg.filePath = path
	cfg.configSet = ConfigSet{}
	err := cfg.Load()
	if err != nil {
		return nil, err
	}

	return cfg, err
}
func (cfg *Config) Load() error {
	err := util.EnsureFile(cfg.filePath)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(cfg.filePath)
	if err != nil {
		return err
	}

	newConfig := ConfigSet{}
	if len(b) > 0 {
		err = json.Unmarshal(b, &newConfig)
		if err != nil {
			return err
		}
	}

	cfg.configSet = newConfig

	return nil
}

func (cfg *Config) Save() error {
	err := util.EnsureFile(cfg.filePath)
	if err != nil {
		return err
	}

	b, err := json.Marshal(cfg.configSet)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(cfg.filePath, b, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) GetLastHeight() string {
	return cfg.configSet.LastHeight
}

func (cfg *Config) SetLastHeight(lastHeight string) {
	cfg.configSet.LastHeight = lastHeight
}

func (cfg *Config) GetABCIConfig() types.AMOAppConfig {
	return cfg.configSet.ABCIConfig
}

func (cfg *Config) SetABCIConfig(abciConfig types.AMOAppConfig) {
	cfg.configSet.ABCIConfig = abciConfig
}

func (cfg *Config) UpdateLastHeight() error {
	rawMsg, err := rpc.NodeStatus()
	if err != nil {
		return err
	}

	if rpc.DryRun {
		return nil
	}

	jsonMsg, err := json.Marshal(rawMsg.SyncInfo)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(jsonMsg, &data)
	if err != nil {
		return err
	}

	lastHeight := data["latest_block_height"].(string)

	if lastHeight == "0" {
		lastHeight = "1"
	}

	cfg.SetLastHeight(lastHeight)
	cfg.Save()

	return nil
}
