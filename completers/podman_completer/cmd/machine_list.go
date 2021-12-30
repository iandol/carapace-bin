package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var machine_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List machines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_listCmd).Standalone()
	machine_listCmd.Flags().String("format", "", "Format volume output using JSON or a Go template")
	machine_listCmd.Flags().Bool("noheading", false, "Do not print headers")
	machineCmd.AddCommand(machine_listCmd)
}
