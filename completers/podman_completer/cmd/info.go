package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display podman system information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().BoolP("debug", "D", false, "Display additional debug information")
	infoCmd.Flags().StringP("format", "f", "", "Change the output format to JSON or a Go template")
	rootCmd.AddCommand(infoCmd)
}
