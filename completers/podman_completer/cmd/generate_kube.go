package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generate_kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Generate Kubernetes YAML from containers, pods or volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_kubeCmd).Standalone()
	generate_kubeCmd.Flags().StringP("filename", "f", "", "Write output to the specified path")
	generate_kubeCmd.Flags().BoolP("service", "s", false, "Generate YAML for a Kubernetes service object")
	generateCmd.AddCommand(generate_kubeCmd)
}
