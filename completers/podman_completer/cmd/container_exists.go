package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Check if a container exists in local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_existsCmd).Standalone()
	container_existsCmd.Flags().Bool("external", false, "Check external storage containers as well as Podman containers")
	containerCmd.AddCommand(container_existsCmd)
}
