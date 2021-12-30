package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_topCmd = &cobra.Command{
	Use:   "top",
	Short: "Display the running processes of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_topCmd).Standalone()
	container_topCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_topCmd.Flags().Bool("list-descriptors", false, "")
	containerCmd.AddCommand(container_topCmd)
}
