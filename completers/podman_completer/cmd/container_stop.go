package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_stopCmd).Standalone()
	container_stopCmd.Flags().BoolP("all", "a", false, "Stop all running containers")
	container_stopCmd.Flags().StringArray("cidfile", []string{}, "Read the container ID from the file")
	container_stopCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified container is missing")
	container_stopCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_stopCmd.Flags().UintP("time", "t", 10, "Seconds to wait for stop before killing the container")
	containerCmd.AddCommand(container_stopCmd)
}
