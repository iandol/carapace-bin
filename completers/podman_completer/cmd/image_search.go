package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search registry for image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_searchCmd).Standalone()
	image_searchCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	image_searchCmd.Flags().Bool("compatible", false, "List stars, official and automated columns (Docker compatibility)")
	image_searchCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions provided (default [])")
	image_searchCmd.Flags().String("format", "", "Change the output format to JSON or a Go template")
	image_searchCmd.Flags().Int("limit", 0, "Limit the number of results")
	image_searchCmd.Flags().Bool("list-tags", false, "List the tags of the input registry")
	image_searchCmd.Flags().Bool("no-trunc", true, "Do not truncate the output. Default: true")
	image_searchCmd.Flags().Bool("tls-verify", true, "Require HTTPS and verify certificates when contacting registries")
	imageCmd.AddCommand(image_searchCmd)
}
