package build

import (
	"log/slog"

	"github.com/alexfalkowski/cli-test/internal/exec"
	"github.com/spf13/cobra"
)

// New build command.
func New(runner exec.Runner, logger *slog.Logger) *cobra.Command {
	build := &cobra.Command{
		Use:   "build",
		Short: "Build with bazel.",
		Long:  "https://bazel.build/run/build",
		RunE: func(c *cobra.Command, args []string) error {
			ctx := c.Context()

			out, err := runner.Run(ctx, "bazel help build")
			if err != nil {
				return err
			}

			logger.InfoContext(ctx, out)

			return nil
		},
	}

	return build
}
