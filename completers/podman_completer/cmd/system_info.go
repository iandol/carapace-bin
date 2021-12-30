package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display podman system information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_infoCmd).Standalone()
	system_infoCmd.Flags().BoolP("debug", "D", false, "Display additional debug information")
	system_infoCmd.Flags().StringP("format", "f", "", "Change the output format to JSON or a Go template")
	systemCmd.AddCommand(system_infoCmd)
}
