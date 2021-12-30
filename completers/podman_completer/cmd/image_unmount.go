package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "Unmount an image's root filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_unmountCmd).Standalone()
	image_unmountCmd.Flags().BoolP("all", "a", false, "Unmount all of the currently mounted images")
	image_unmountCmd.Flags().BoolP("force", "f", false, "Force the complete unmount of the specified mounted images")
	imageCmd.AddCommand(image_unmountCmd)
}
