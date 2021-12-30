package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var volume_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_createCmd).Standalone()
	volume_createCmd.Flags().String("driver", "local", "Specify volume driver name")
	volume_createCmd.Flags().StringSliceP("label", "l", []string{}, "Set metadata for a volume (default [])")
	volume_createCmd.Flags().StringArrayP("opt", "o", []string{}, "Set driver specific options (default [])")
	volumeCmd.AddCommand(volume_createCmd)
}
