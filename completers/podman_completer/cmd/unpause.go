package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause the processes in one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpauseCmd).Standalone()
	unpauseCmd.Flags().BoolP("all", "a", false, "Pause all running containers")
	rootCmd.AddCommand(unpauseCmd)
}
