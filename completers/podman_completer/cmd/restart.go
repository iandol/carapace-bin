package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()
	restartCmd.Flags().BoolP("all", "a", false, "Restart all non-running containers")
	restartCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	restartCmd.Flags().Bool("running", false, "Restart only running containers when --all is used")
	restartCmd.Flags().UintP("time", "t", 10, "Seconds to wait for stop before killing the container")
	rootCmd.AddCommand(restartCmd)
}
