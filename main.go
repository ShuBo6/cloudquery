package main

import (
	"os"

	"github.com/ShuBo6/cloudquery/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
