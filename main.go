package main

import (
	"os"

	"github.com/Dedalum/goatter/cli"
)

func main() {
	defer os.Exit(0)

	cli := cli.CommandLine{}

	cli.Run()

}
