package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup network and mountpoints of one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cleanupCmd).Standalone()
	container_cleanupCmd.Flags().BoolP("all", "a", false, "Cleans up all containers")
	container_cleanupCmd.Flags().String("exec", "", "Clean up the given exec session instead of the container")
	container_cleanupCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_cleanupCmd.Flags().Bool("rm", false, "After cleanup, remove the container entirely")
	container_cleanupCmd.Flags().Bool("rmi", false, "After cleanup, remove the image entirely")
	containerCmd.AddCommand(container_cleanupCmd)
}
