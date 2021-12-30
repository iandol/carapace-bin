package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_restartCmd).Standalone()
	container_restartCmd.Flags().BoolP("all", "a", false, "Restart all non-running containers")
	container_restartCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_restartCmd.Flags().Bool("running", false, "Restart only running containers when --all is used")
	container_restartCmd.Flags().UintP("time", "t", 10, "Seconds to wait for stop before killing the container")
	containerCmd.AddCommand(container_restartCmd)
}
