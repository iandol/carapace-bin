package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Displays the raw CNI network configuration for one or more networks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_inspectCmd).Standalone()
	network_inspectCmd.Flags().StringP("format", "f", "", "Pretty-print network to JSON or using a Go template")
	networkCmd.AddCommand(network_inspectCmd)
}
