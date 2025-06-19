package build

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// New build command.
func New() *cobra.Command {
	build := &cobra.Command{
		Use:   "build",
		Short: "Build with bazel.",
		Long:  "https://bazel.build/run/build",
		RunE: func(c *cobra.Command, args []string) error {
			cmd := exec.CommandContext(c.Context(), "bazel", "help", "build")

			data, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}

			fmt.Println(string(data))

			return nil
		},
	}

	return build
}
