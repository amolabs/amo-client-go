package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	defaultCLIDir      = ".amocli"
	defaultConfigFile  = "config.json"
	defaultKeyDir      = "keys"
	defaultKeyListFile = "keys.json"
)

func defaultCLIPath() string {
	return filepath.Join(os.ExpandEnv("$HOME"), defaultCLIDir)
}

func DefaultConfigFilePath() string {
	return filepath.Join(defaultCLIPath(), defaultConfigFile)
}

func DefaultKeyPath() string {
	return filepath.Join(defaultCLIPath(), defaultKeyDir)
}

func DefaultKeyFilePath() string {
	return filepath.Join(DefaultKeyPath(), defaultKeyListFile)
}

func PromptUsername() (string, error) {
	fmt.Printf("\nInput username of the signing key: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(username, "\r\n"), nil
}

func PromptPassphrase() (string, error) {
	fmt.Printf("Type passphrase: ")
	b, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var LineBreak = &cobra.Command{Run: func(*cobra.Command, []string) {}}

func EnsureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}

func EnsureFile(path string) error {
	dirPath, _ := filepath.Split(path)

	if len(dirPath) > 0 {
		err := EnsureDir(dirPath, 0775)
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
