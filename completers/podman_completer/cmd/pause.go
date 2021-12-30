package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause all the processes in one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pauseCmd).Standalone()
	pauseCmd.Flags().BoolP("all", "a", false, "Pause all running containers")
	rootCmd.AddCommand(pauseCmd)
}
