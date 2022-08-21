package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "dink <command>",
	Short:         "Dink Docs - keep your docs in sync",
	Example:       "",
	SilenceUsage:  false,
	SilenceErrors: false,
}

func init() {
	rootCmd.Version = Version
	rootCmd.Flags().BoolP("version", "v", false, "Get the version of current Dink CLI")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "Update the docs to dependencies snapshot",
		RunE:  UpdateCmd,
	})
}

func main() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
	}
}
