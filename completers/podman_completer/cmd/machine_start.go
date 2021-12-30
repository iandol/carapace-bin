package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var machine_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_startCmd).Standalone()
	machineCmd.AddCommand(machine_startCmd)
}
