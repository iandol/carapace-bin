package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var machine_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a virtual machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_initCmd).Standalone()
	machine_initCmd.Flags().Uint64("cpus", 1, "Number of CPUs")
	machine_initCmd.Flags().Uint64("disk-size", 100, "Disk size in GB")
	machine_initCmd.Flags().String("ignition-path", "", "Path to ignition file")
	machine_initCmd.Flags().String("image-path", "testing", "Path to qcow image")
	machine_initCmd.Flags().Uint64P("memory", "m", 2048, "Memory in MB")
	machine_initCmd.Flags().Bool("now", false, "Start machine now")
	machine_initCmd.Flags().Bool("reexec", false, "process was rexeced")
	machine_initCmd.Flags().String("timezone", "local", "Set timezone")
	machineCmd.AddCommand(machine_initCmd)
}
