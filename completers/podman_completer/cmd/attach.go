package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()
	attachCmd.Flags().String("detach-keys", "ctrl-p,ctrl-q", "Select the key sequence for detaching a container.")
	attachCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	attachCmd.Flags().Bool("no-stdin", false, "Do not attach STDIN. The default is false")
	attachCmd.Flags().Bool("sig-proxy", true, "Proxy received signals to the process")
	rootCmd.AddCommand(attachCmd)
}
