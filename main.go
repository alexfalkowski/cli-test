package main

import (
	"log/slog"
	"os"

	"github.com/alexfalkowski/cli-test/internal/cmd/build"
	"github.com/alexfalkowski/cli-test/internal/cmd/root"
	"github.com/alexfalkowski/cli-test/internal/exec"
	"github.com/alexfalkowski/cli-test/internal/exec/logger"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	runner := runner(logger)
	root := root.New()
	root.AddCommand(build.New(runner, logger))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func dryRun() bool {
	d := os.Getenv("DRY_RUN")

	return d == "true" || d == "1"
}

func runner(l *slog.Logger) exec.Runner {
	if dryRun() {
		return logger.NewCommand(l)
	}

	return exec.NewCommand()
}
