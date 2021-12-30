package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill one or more running containers with a specific signal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_killCmd).Standalone()
	container_killCmd.Flags().BoolP("all", "a", false, "Signal all running containers")
	container_killCmd.Flags().StringArray("cidfile", []string{}, "Read the container ID from the file")
	container_killCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_killCmd.Flags().StringP("signal", "s", "KILL", "Signal to send to the container")
	containerCmd.AddCommand(container_killCmd)
}
