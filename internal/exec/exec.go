package exec

import (
	"context"
	"os/exec"
	"strings"
)

var _ Runner = (*Command)(nil)

// NewCommand returns a command using os/exec.
func NewCommand() *Command {
	return &Command{}
}

// CommandRunner uses os/exec.
type Command struct{}

// Run cmd and return the output, unless an error occurs.
func (r *Command) Run(ctx context.Context, args string) (string, error) {
	split := strings.Split(args, " ")
	cmd := exec.CommandContext(ctx, split[0], split[1:]...)

	data, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(data), nil
}
