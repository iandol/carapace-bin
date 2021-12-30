package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Run API service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_serviceCmd).Standalone()
	system_serviceCmd.Flags().String("cors", "", "Set CORS Headers")
	system_serviceCmd.Flags().String("pprof-address", "", "Binding network address for pprof profile endpoints, default: do not expose endpoints")
	system_serviceCmd.Flags().UintP("time", "t", 5, "Time until the service session expires in seconds.  Use 0 to disable the timeout")
	systemCmd.AddCommand(system_serviceCmd)
}
