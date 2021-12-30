package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "network rm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_disconnectCmd).Standalone()
	network_disconnectCmd.Flags().BoolP("force", "f", false, "force removal of container from network")
	networkCmd.AddCommand(network_disconnectCmd)
}
