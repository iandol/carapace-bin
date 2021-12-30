package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pod_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_startCmd).Standalone()
	pod_startCmd.Flags().BoolP("all", "a", false, "Restart all running pods")
	pod_startCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_startCmd.Flags().StringArray("pod-id-file", []string{}, "Read the pod ID from the file")
	podCmd.AddCommand(pod_startCmd)
}
