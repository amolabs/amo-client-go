package main

import (
	"fmt"
	"os"

	"github.com/amolabs/amo-client-go/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
