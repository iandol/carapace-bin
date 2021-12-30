package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_connection_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Record destination for the Podman service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_connection_addCmd).Standalone()
	system_connection_addCmd.Flags().BoolP("default", "d", false, "Set connection to be default")
	system_connection_addCmd.Flags().String("identity", "", "path to SSH identity file")
	system_connection_addCmd.Flags().IntP("port", "p", 22, "SSH port number for destination")
	system_connection_addCmd.Flags().String("socket-path", "", "path to podman socket on remote host. (default '/run/podman/podman.sock' or '/run/user/{uid}/podman/podman.sock)")
	system_connectionCmd.AddCommand(system_connection_addCmd)
}
