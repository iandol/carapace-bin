package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_untagCmd = &cobra.Command{
	Use:   "untag",
	Short: "Remove a name from a local image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_untagCmd).Standalone()
	imageCmd.AddCommand(image_untagCmd)
}
