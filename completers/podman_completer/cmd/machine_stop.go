package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var machine_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_stopCmd).Standalone()
	machineCmd.AddCommand(machine_stopCmd)
}
