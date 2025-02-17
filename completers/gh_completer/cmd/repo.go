package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Create, clone, fork, and view repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	rootCmd.AddCommand(repoCmd)
}
