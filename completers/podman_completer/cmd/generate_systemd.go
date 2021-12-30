package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generate_systemdCmd = &cobra.Command{
	Use:   "systemd",
	Short: "Generate systemd units.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_systemdCmd).Standalone()
	generate_systemdCmd.Flags().String("container-prefix", "container", "Systemd unit name prefix for containers")
	generate_systemdCmd.Flags().BoolP("files", "f", false, "Generate .service files instead of printing to stdout")
	generate_systemdCmd.Flags().String("format", "", "Print the created units in specified format (json)")
	generate_systemdCmd.Flags().BoolP("name", "n", false, "Use container/pod names instead of IDs")
	generate_systemdCmd.Flags().Bool("new", false, "Create a new container or pod instead of starting an existing one")
	generate_systemdCmd.Flags().Bool("no-header", false, "Skip header generation")
	generate_systemdCmd.Flags().String("pod-prefix", "pod", "Systemd unit name prefix for pods")
	generate_systemdCmd.Flags().String("restart-policy", "on-failure", "Systemd restart-policy")
	generate_systemdCmd.Flags().Uint("restart-sec", 0, "Systemd restart-sec")
	generate_systemdCmd.Flags().String("separator", "-", "Systemd unit name separator between name/id and prefix")
	generate_systemdCmd.Flags().Uint("start-timeout", 0, "Start timeout override")
	generate_systemdCmd.Flags().Uint("stop-timeout", 10, "Stop timeout override")
	generate_systemdCmd.Flags().Bool("template", false, "Make it a template file and use %i and %I specifiers. Working only for containers")
	generate_systemdCmd.Flags().UintP("time", "t", 10, "Backwards alias for --stop-timeout")
	generateCmd.AddCommand(generate_systemdCmd)
}
