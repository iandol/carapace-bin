package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files/folders between a container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()
	cpCmd.Flags().BoolP("archive", "a", true, "Chown copied files to the primary uid/gid of the destination container.")
	cpCmd.Flags().Bool("extract", false, "Deprecated...")
	cpCmd.Flags().Bool("pause", true, "Deprecated")
	rootCmd.AddCommand(cpCmd)
}
