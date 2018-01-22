package main

import (
	"os"

	"github.com/DuncanQ/lsbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
