package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var volume_existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "volume exists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_existsCmd).Standalone()
	volumeCmd.AddCommand(volume_existsCmd)
}
