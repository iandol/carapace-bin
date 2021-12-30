package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an existing container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_renameCmd).Standalone()
	containerCmd.AddCommand(container_renameCmd)
}
