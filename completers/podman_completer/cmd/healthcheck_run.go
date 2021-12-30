package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var healthcheck_runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the health check of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthcheck_runCmd).Standalone()
	healthcheckCmd.AddCommand(healthcheck_runCmd)
}
