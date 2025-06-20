package logger

import (
	"context"
	"log/slog"

	"github.com/alexfalkowski/cli-test/internal/exec"
)

var _ exec.Runner = (*Command)(nil)

// NewCommand returns a command that just logs.
func NewCommand(logger *slog.Logger) *Command {
	return &Command{logger: logger}
}

// Command that logs.
type Command struct {
	logger *slog.Logger
}

// Run just logs the command.
func (r *Command) Run(ctx context.Context, args string) (string, error) {
	r.logger.InfoContext(ctx, "running command", slog.String("cmd", args))

	return "", nil
}
