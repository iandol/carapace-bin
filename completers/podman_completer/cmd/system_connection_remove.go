package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_connection_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete named destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_removeCmd).Standalone()
	system_connection_removeCmd.Flags().BoolP("all", "a", false, "Remove all connections")
	system_connectionCmd.AddCommand(system_connection_removeCmd)
}
