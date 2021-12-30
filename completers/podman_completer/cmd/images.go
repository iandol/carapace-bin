package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images in local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagesCmd).Standalone()
	imagesCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	imagesCmd.Flags().Bool("digests", false, "Show digests")
	imagesCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions provided (default [])")
	imagesCmd.Flags().String("format", "", "Change the output format to JSON or a Go template")
	imagesCmd.Flags().Bool("history", false, "Display the image name history")
	imagesCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	imagesCmd.Flags().BoolP("noheading", "n", false, "Do not print column headings")
	imagesCmd.Flags().BoolP("quiet", "q", false, "Display only image IDs")
	imagesCmd.Flags().String("sort", "created", "Sort by created, id, repository, size, tag")
	rootCmd.AddCommand(imagesCmd)
}
