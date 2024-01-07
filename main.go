package main

import (
	"os"
	"github.com/btwkenji/blobman/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
