package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// buildCmd represents the serve command
var buildCmd = &cobra.Command{
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

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
