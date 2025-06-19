package root

import "github.com/spf13/cobra"

// New root command.
func New() *cobra.Command {
	return &cobra.Command{
		Use:   "cli-test",
		Short: "A test cli.",
		Long:  "Demo how to test CLIs.",
	}
}
