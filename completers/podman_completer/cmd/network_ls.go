package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "network list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_lsCmd).Standalone()
	network_lsCmd.Flags().StringArrayP("filter", "f", []string{}, "Provide filter values (e.g. 'name=podman')")
	network_lsCmd.Flags().String("format", "", "Pretty-print networks to JSON or using a Go template")
	network_lsCmd.Flags().Bool("no-trunc", false, "Do not truncate the network ID")
	network_lsCmd.Flags().Bool("noheading", false, "Do not print headers")
	network_lsCmd.Flags().BoolP("quiet", "q", false, "display only names")
	networkCmd.AddCommand(network_lsCmd)
}
