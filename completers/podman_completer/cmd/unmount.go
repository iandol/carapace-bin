package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "Unmounts working container's root filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmountCmd).Standalone()
	unmountCmd.Flags().BoolP("all", "a", false, "Unmount all of the currently mounted containers")
	unmountCmd.Flags().BoolP("force", "f", false, "Force the complete unmount of the specified mounted containers")
	unmountCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(unmountCmd)
}
