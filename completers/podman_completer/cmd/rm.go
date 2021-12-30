package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()
	rmCmd.Flags().BoolP("all", "a", false, "Remove all containers")
	rmCmd.Flags().StringArray("cidfile", []string{}, "Read the container ID from the file")
	rmCmd.Flags().BoolP("force", "f", false, "Force removal of a running or unusable container.  The default is false")
	rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified container is missing")
	rmCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rmCmd.Flags().Bool("storage", false, "Remove container from storage library")
	rmCmd.Flags().UintP("time", "t", 10, "Seconds to wait for stop before killing the container")
	rmCmd.Flags().BoolP("volumes", "v", false, "Remove anonymous volumes associated with the container")
	rootCmd.AddCommand(rmCmd)
}
