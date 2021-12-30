package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_createCmd).Standalone()
	manifest_createCmd.Flags().Bool("all", false, "add all of the lists' images if the images to add are lists")
	manifestCmd.AddCommand(manifest_createCmd)
}
