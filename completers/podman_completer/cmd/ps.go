package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(psCmd).Standalone()
	psCmd.Flags().BoolP("all", "a", false, "Show all the containers, default is only running containers")
	psCmd.Flags().Bool("external", false, "Show containers in storage not controlled by Podman")
	psCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	psCmd.Flags().String("format", "", "Pretty-print containers to JSON or using a Go template")
	psCmd.Flags().IntP("last", "n", -1, "Print the n last created containers (all states)")
	psCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	psCmd.Flags().Bool("no-trunc", false, "Display the extended information")
	psCmd.Flags().Bool("noheading", false, "Do not print headers")
	psCmd.Flags().BoolP("pod", "p", false, "Print the ID and name of the pod the containers are associated with")
	psCmd.Flags().BoolP("quiet", "q", false, "Print the numeric IDs of the containers only")
	psCmd.Flags().BoolP("size", "s", false, "Display the total file sizes")
	psCmd.Flags().String("sort", "", "Sort output by: command, created, id, image, names, runningfor, size, status")
	psCmd.Flags().Bool("sync", false, "Sync container state with OCI runtime")
	psCmd.Flags().UintP("watch", "w", 0, "Watch the ps output on an interval in seconds")
	rootCmd.AddCommand(psCmd)
}
