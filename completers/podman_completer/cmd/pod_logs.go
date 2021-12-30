package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pod_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch logs for pod with one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_logsCmd).Standalone()
	pod_logsCmd.Flags().StringP("container", "c", "", "Filter logs by container name or id which belongs to pod")
	pod_logsCmd.Flags().Bool("details", false, "Show extra details provided to the logs")
	pod_logsCmd.Flags().BoolP("follow", "f", false, "Follow log output.")
	pod_logsCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_logsCmd.Flags().String("since", "", "Show logs since TIMESTAMP")
	pod_logsCmd.Flags().Int64("tail", -1, "Output the specified number of LINES at the end of the logs.")
	pod_logsCmd.Flags().BoolP("timestamps", "t", false, "Output the timestamps in the log")
	pod_logsCmd.Flags().String("until", "", "Show logs until TIMESTAMP")
	podCmd.AddCommand(pod_logsCmd)
}
