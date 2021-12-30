package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "network rm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rmCmd).Standalone()
	network_rmCmd.Flags().BoolP("force", "f", false, "remove any containers using network")
	network_rmCmd.Flags().UintP("time", "t", 10, "Seconds to wait for running containers to stop before killing the container")
	networkCmd.AddCommand(network_rmCmd)
}
