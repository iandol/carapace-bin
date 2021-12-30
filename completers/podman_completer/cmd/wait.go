package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Block on one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()
	waitCmd.Flags().String("condition", "stopped", "Condition to wait on")
	waitCmd.Flags().StringP("interval", "i", "250ms", "Time Interval to wait before polling for completion")
	waitCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(waitCmd)
}
