package util

import (
	"os"
	"path/filepath"

	cmn "github.com/tendermint/tendermint/libs/common"
)

func EnsureFile(path string) error {
	dirPath, _ := filepath.Split(path)

	if len(dirPath) > 0 {
		err := cmn.EnsureDir(dirPath, 0775)
		if err != nil {
			return err
		}
	}

	_, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	return err
}
