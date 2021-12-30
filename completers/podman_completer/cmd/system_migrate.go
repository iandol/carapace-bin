package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_migrateCmd).Standalone()
	system_migrateCmd.Flags().String("new-runtime", "", "Specify a new runtime for all containers")
	systemCmd.AddCommand(system_migrateCmd)
}
