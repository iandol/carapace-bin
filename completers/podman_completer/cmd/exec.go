package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a process in a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()
	execCmd.Flags().BoolP("detach", "d", false, "Run the exec session in detached mode (backgrounded)")
	execCmd.Flags().String("detach-keys", "ctrl-p,ctrl-q", "Select the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _")
	execCmd.Flags().StringArrayP("env", "e", []string{}, "Set environment variables")
	execCmd.Flags().StringSlice("env-file", []string{}, "Read in a file of environment variables")
	execCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	execCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	execCmd.Flags().Uint("preserve-fds", 0, "Pass N additional file descriptors to the container")
	execCmd.Flags().Bool("privileged", false, "Give the process extended Linux capabilities inside the container.  The default is false")
	execCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY. The default is false")
	execCmd.Flags().StringP("user", "u", "", "Sets the username or UID used and optionally the groupname or GID for the specified command")
	execCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	rootCmd.AddCommand(execCmd)
}
