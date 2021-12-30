package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Add or update information about an entry in a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_annotateCmd).Standalone()
	manifest_annotateCmd.Flags().StringSlice("annotation", []string{}, "set an `annotation` for the specified image")
	manifest_annotateCmd.Flags().String("arch", "", "override the `architecture` of the specified image")
	manifest_annotateCmd.Flags().StringSlice("features", []string{}, "override the `features` of the specified image")
	manifest_annotateCmd.Flags().String("os", "", "override the `OS` of the specified image")
	manifest_annotateCmd.Flags().StringSlice("os-features", []string{}, "override the OS `features` of the specified image")
	manifest_annotateCmd.Flags().String("os-version", "", "override the OS `version` of the specified image")
	manifest_annotateCmd.Flags().String("variant", "", "override the `Variant` of the specified image")
	manifestCmd.AddCommand(manifest_annotateCmd)
}
