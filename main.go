package main

import (
	"os"

	"github.com/airbowen/golang-blockchain/blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
