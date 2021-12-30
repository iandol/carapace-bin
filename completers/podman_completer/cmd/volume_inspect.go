package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var volume_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_inspectCmd).Standalone()
	volume_inspectCmd.Flags().BoolP("all", "a", false, "Inspect all volumes")
	volume_inspectCmd.Flags().StringP("format", "f", "json", "Format volume output using Go template")
	volumeCmd.AddCommand(volume_inspectCmd)
}
