package main

import (
	"log/slog"
	"os"

	"github.com/alexfalkowski/cli-test/internal/cmd/build"
	"github.com/alexfalkowski/cli-test/internal/cmd/root"
	"github.com/alexfalkowski/cli-test/internal/exec"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	runner := exec.NewCommand()
	root := root.New()
	root.AddCommand(build.New(runner, logger))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
