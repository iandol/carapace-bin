package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Inspect changes to the image's file systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_diffCmd).Standalone()
	image_diffCmd.Flags().Bool("archive", true, "Save the diff as a tar archive")
	image_diffCmd.Flags().String("format", "", "Change the output format (json)")
	imageCmd.AddCommand(image_diffCmd)
}
