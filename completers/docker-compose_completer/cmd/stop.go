package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()
	stopCmd.Flags().IntP("timeout", "t", 10, "Specify a shutdown timeout in seconds")
	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).PositionalAnyCompletion(
		action.ActionServices(stopCmd),
	)
}
