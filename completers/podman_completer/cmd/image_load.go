package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load image(s) from a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_loadCmd).Standalone()
	image_loadCmd.Flags().StringP("input", "i", "", "Read from specified archive file (default: stdin)")
	image_loadCmd.Flags().BoolP("quiet", "q", false, "Suppress the output")
	image_loadCmd.Flags().String("signature-policy", "", "Pathname of signature policy file")
	imageCmd.AddCommand(image_loadCmd)
}
