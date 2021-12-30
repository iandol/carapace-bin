package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display the configuration of an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_inspectCmd).Standalone()
	image_inspectCmd.Flags().StringP("format", "f", "json", "Format the output to a Go template or json")
	imageCmd.AddCommand(image_inspectCmd)
}
