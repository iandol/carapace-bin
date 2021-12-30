package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Display the running processes of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(topCmd).Standalone()
	topCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	topCmd.Flags().Bool("list-descriptors", false, "")
	rootCmd.AddCommand(topCmd)
}
