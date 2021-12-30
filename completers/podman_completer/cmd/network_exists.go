package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "network exists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_existsCmd).Standalone()
	networkCmd.AddCommand(network_existsCmd)
}
