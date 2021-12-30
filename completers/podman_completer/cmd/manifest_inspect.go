package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display the contents of a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_inspectCmd).Standalone()
	manifestCmd.AddCommand(manifest_inspectCmd)
}
