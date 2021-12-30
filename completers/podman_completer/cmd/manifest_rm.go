package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove manifest list or image index from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_rmCmd).Standalone()
	manifestCmd.AddCommand(manifest_rmCmd)
}
