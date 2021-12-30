package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()
	startCmd.Flags().Bool("all", false, "Start all containers regardless of their state or configuration")
	startCmd.Flags().BoolP("attach", "a", false, "Attach container's STDOUT and STDERR")
	startCmd.Flags().String("detach-keys", "ctrl-p,ctrl-q", "Select the key sequence for detaching a container.")
	startCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	startCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	startCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	startCmd.Flags().Bool("sig-proxy", false, "Proxy received signals to the process (default true if attaching, false otherwise)")
	rootCmd.AddCommand(startCmd)
}
