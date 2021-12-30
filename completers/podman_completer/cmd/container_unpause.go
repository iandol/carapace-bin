package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause the processes in one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_unpauseCmd).Standalone()
	container_unpauseCmd.Flags().BoolP("all", "a", false, "Pause all running containers")
	containerCmd.AddCommand(container_unpauseCmd)
}
