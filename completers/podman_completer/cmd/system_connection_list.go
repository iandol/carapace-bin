package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_connection_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List destination for the Podman service(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_listCmd).Standalone()
	system_connection_listCmd.Flags().String("format", "", "Custom Go template for printing connections")
	system_connectionCmd.AddCommand(system_connection_listCmd)
}
