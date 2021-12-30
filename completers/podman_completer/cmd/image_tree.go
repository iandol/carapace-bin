package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Prints layer hierarchy of an image in a tree format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_treeCmd).Standalone()
	image_treeCmd.Flags().Bool("whatrequires", false, "Show all child images and layers of the specified image")
	imageCmd.AddCommand(image_treeCmd)
}
