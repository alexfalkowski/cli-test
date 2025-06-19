package build

import (
	"fmt"

	"github.com/alexfalkowski/cli-test/internal/exec"
	"github.com/spf13/cobra"
)

// New build command.
func New(runner exec.Runner) *cobra.Command {
	build := &cobra.Command{
		Use:   "build",
		Short: "Build with bazel.",
		Long:  "https://bazel.build/run/build",
		RunE: func(c *cobra.Command, args []string) error {
			data, err := runner.Run(c.Context(), "bazel help build")
			if err != nil {
				return err
			}

			fmt.Println(string(data))

			return nil
		},
	}

	return build
}
