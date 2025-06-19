package test

import (
	"context"

	"github.com/alexfalkowski/cli-test/internal/exec"
)

// Outputs allows us to control what is returned by which command.
type Outputs map[string]any

var _ exec.Runner = (*Command)(nil)

// NewCommand returns a command for testing.
func NewCommand(outputs Outputs) *Command {
	return &Command{
		outputs: outputs,
		Cmds:    []string{},
	}
}

// Command for test.
type Command struct {
	outputs Outputs
	Cmds    []string
}

// Run appends the commands and returns what was passed in from outputs.
func (r *Command) Run(ctx context.Context, args string) (string, error) {
	r.Cmds = append(r.Cmds, args)

	out, ok := r.outputs[args]
	if ok {
		switch v := out.(type) {
		case string:
			return v, nil
		case error:
			return "", v
		}
	}

	return "", nil
}
