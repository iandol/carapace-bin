package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an entry from a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_removeCmd).Standalone()
	manifestCmd.AddCommand(manifest_removeCmd)
}
