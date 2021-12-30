package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search registry for image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	searchCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	searchCmd.Flags().Bool("compatible", false, "List stars, official and automated columns (Docker compatibility)")
	searchCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions provided (default [])")
	searchCmd.Flags().String("format", "", "Change the output format to JSON or a Go template")
	searchCmd.Flags().Int("limit", 0, "Limit the number of results")
	searchCmd.Flags().Bool("list-tags", false, "List the tags of the input registry")
	searchCmd.Flags().Bool("no-trunc", true, "Do not truncate the output. Default: true")
	searchCmd.Flags().Bool("tls-verify", true, "Require HTTPS and verify certificates when contacting registries")
	rootCmd.AddCommand(searchCmd)
}
