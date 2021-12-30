package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Display the changes to the object's file system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	diffCmd.Flags().Bool("archive", true, "Save the diff as a tar archive")
	diffCmd.Flags().String("format", "", "Change the output format (json)")
	diffCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(diffCmd)
}
