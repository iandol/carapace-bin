package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an existing container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameCmd).Standalone()
	rootCmd.AddCommand(renameCmd)
}
