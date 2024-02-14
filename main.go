package main

import (
	"os"
	"github.com/kenjitheman/blobman/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
