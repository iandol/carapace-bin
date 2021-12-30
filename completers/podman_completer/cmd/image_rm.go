package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes one or more images from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_rmCmd).Standalone()
	image_rmCmd.Flags().BoolP("all", "a", false, "Remove all images")
	image_rmCmd.Flags().BoolP("force", "f", false, "Force Removal of the image")
	imageCmd.AddCommand(image_rmCmd)
}
