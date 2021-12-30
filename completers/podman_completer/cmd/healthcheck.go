package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Manage health checks on containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthcheckCmd).Standalone()
	rootCmd.AddCommand(healthcheckCmd)
}
