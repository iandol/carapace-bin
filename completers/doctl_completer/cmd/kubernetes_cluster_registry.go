package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Display commands for integrating clusters with docr",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_registryCmd).Standalone()
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_registryCmd)
}
