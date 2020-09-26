package main

import (
	"os"

	"github.com/wborbajr/gochain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
