package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pod_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Displays a pod configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_inspectCmd).Standalone()
	pod_inspectCmd.Flags().StringP("format", "f", "json", "Format the output to a Go template or json")
	pod_inspectCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	podCmd.AddCommand(pod_inspectCmd)
}
