package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files/folders between a container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cpCmd).Standalone()
	container_cpCmd.Flags().BoolP("archive", "a", true, "Chown copied files to the primary uid/gid of the destination container.")
	container_cpCmd.Flags().Bool("extract", false, "Deprecated...")
	container_cpCmd.Flags().Bool("pause", true, "Deprecated")
	containerCmd.AddCommand(container_cpCmd)
}
