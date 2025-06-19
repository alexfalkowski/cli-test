package build_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/alexfalkowski/cli-test/internal/cmd/build"
	"github.com/alexfalkowski/cli-test/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestBuildSuccess(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	outputs := test.Outputs{
		"bazel help build": "test",
	}
	runner := test.NewCommand(outputs)

	cmd := build.New(runner, logger)
	err := cmd.Execute()

	assert.NoError(t, err)
	assert.Len(t, runner.Cmds, 1)
	assert.Equal(t, "bazel help build", runner.Cmds[0])
}

func TestBuildError(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	outputs := test.Outputs{
		"bazel help build": errors.New("test"),
	}
	runner := test.NewCommand(outputs)

	cmd := build.New(runner, logger)
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Len(t, runner.Cmds, 1)
	assert.Equal(t, "bazel help build", runner.Cmds[0])
}
