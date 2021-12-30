package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var machine_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_rmCmd).Standalone()
	machine_rmCmd.Flags().Bool("force", false, "Do not prompt before rming")
	machine_rmCmd.Flags().Bool("save-ignition", false, "Do not delete ignition file")
	machine_rmCmd.Flags().Bool("save-image", false, "Do not delete the image file")
	machine_rmCmd.Flags().Bool("save-keys", false, "Do not delete SSH keys")
	machineCmd.AddCommand(machine_rmCmd)
}
