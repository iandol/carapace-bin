package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "generate the autocompletion script for the specified shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()
	rootCmd.AddCommand(completionCmd)
}
