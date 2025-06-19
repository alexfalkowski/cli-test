package main

import (
	"os"

	"github.com/alexfalkowski/cli-test/internal/cmd/build"
	"github.com/alexfalkowski/cli-test/internal/cmd/root"
)

func main() {
	root := root.New()
	root.AddCommand(build.New())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
