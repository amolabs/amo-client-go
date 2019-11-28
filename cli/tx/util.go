package tx

import (
	"github.com/amolabs/amo-client-go/lib/config"
)

func GetLastHeight(path string) (string, error) {
	lastHeight := ""
	cfg, err := config.GetConfig(path)
	if err != nil {
		return lastHeight, err
	}

	err = cfg.UpdateLastHeight()
	if err != nil {
		return lastHeight, err
	}

	return cfg.GetLastHeight(), nil
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
