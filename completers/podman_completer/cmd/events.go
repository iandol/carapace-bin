package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Show podman events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()
	eventsCmd.Flags().StringArray("filter", []string{}, "filter output")
	eventsCmd.Flags().String("format", "", "format the output using a Go template")
	eventsCmd.Flags().Bool("no-trunc", true, "do not truncate the output")
	eventsCmd.Flags().String("since", "", "show all events created since timestamp")
	eventsCmd.Flags().Bool("stream", true, "stream new events; for testing only")
	eventsCmd.Flags().String("until", "", "show all events until timestamp")
	rootCmd.AddCommand(eventsCmd)
}
