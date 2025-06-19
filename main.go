package main

import (
	"os"

	"github.com/alexfalkowski/cli-test/internal/cmd/build"
	"github.com/alexfalkowski/cli-test/internal/cmd/root"
	"github.com/alexfalkowski/cli-test/internal/exec"
)

func main() {
	runner := exec.NewCommand()
	root := root.New()
	root.AddCommand(build.New(runner))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
