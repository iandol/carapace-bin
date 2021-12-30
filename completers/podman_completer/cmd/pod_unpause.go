package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pod_unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_unpauseCmd).Standalone()
	pod_unpauseCmd.Flags().BoolP("all", "a", false, "Pause all running pods")
	pod_unpauseCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	podCmd.AddCommand(pod_unpauseCmd)
}
