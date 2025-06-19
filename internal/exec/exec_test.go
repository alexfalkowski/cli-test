package exec_test

import (
	"testing"

	"github.com/alexfalkowski/cli-test/internal/exec"
	"github.com/stretchr/testify/assert"
)

func TestRunSuccess(t *testing.T) {
	runner := exec.NewCommand()
	out, err := runner.Run(t.Context(), "date")

	assert.NoError(t, err)
	assert.NotEmpty(t, out)
}

func TestRunError(t *testing.T) {
	runner := exec.NewCommand()
	_, err := runner.Run(t.Context(), "bob")

	assert.Error(t, err)
}
