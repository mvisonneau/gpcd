package main

import (
	"os"

	"github.com/mvisonneau/gpcd/internal/cli"
)

var version = ""

func main() {
	cli.Run(version, os.Args)
}
