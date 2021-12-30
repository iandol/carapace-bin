package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "Removes one or more images from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmiCmd).Standalone()
	rmiCmd.Flags().BoolP("all", "a", false, "Remove all images")
	rmiCmd.Flags().BoolP("force", "f", false, "Force Removal of the image")
	rootCmd.AddCommand(rmiCmd)
}
