package build_test

import (
	"errors"
	"testing"

	"github.com/alexfalkowski/cli-test/internal/cmd/build"
	"github.com/alexfalkowski/cli-test/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestBuildSuccess(t *testing.T) {
	outputs := test.Outputs{
		"bazel help build": "test",
	}
	runner := test.NewCommand(outputs)

	cmd := build.New(runner)
	err := cmd.Execute()

	assert.NoError(t, err)
	assert.Equal(t, "bazel help build", runner.Cmds[0])
}

func TestBuildError(t *testing.T) {
	outputs := test.Outputs{
		"bazel help build": errors.New("test"),
	}
	runner := test.NewCommand(outputs)

	cmd := build.New(runner)
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Equal(t, "bazel help build", runner.Cmds[0])
}
