package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add an additional name to a local image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()
	rootCmd.AddCommand(tagCmd)
}
