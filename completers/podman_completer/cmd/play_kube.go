package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var play_kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Play a pod or volume based on Kubernetes YAML.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(play_kubeCmd).Standalone()
	play_kubeCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	play_kubeCmd.Flags().Bool("build", false, "Build all images in a YAML (given Containerfiles exist)")
	play_kubeCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	play_kubeCmd.Flags().StringSlice("configmap", []string{}, "`Pathname` of a YAML file containing a kubernetes configmap")
	play_kubeCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	play_kubeCmd.Flags().Bool("down", false, "Stop pods defined in the YAML file")
	play_kubeCmd.Flags().StringArray("ip", []string{}, "Static IP addresses to assign to the pods")
	play_kubeCmd.Flags().String("log-driver", "", "Logging driver for the container")
	play_kubeCmd.Flags().StringSlice("log-opt", []string{}, "Logging driver options")
	play_kubeCmd.Flags().StringSlice("mac-address", []string{}, "Static MAC addresses to assign to the pods")
	play_kubeCmd.Flags().StringArray("network", []string{}, "Connect pod to network(s) or network mode")
	play_kubeCmd.Flags().Bool("no-hosts", false, "Do not create /etc/hosts within the pod's containers, instead use the version from the image")
	play_kubeCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pulling images")
	play_kubeCmd.Flags().Bool("replace", false, "Delete and recreate pods defined in the YAML file")
	play_kubeCmd.Flags().String("seccomp-profile-root", "/var/lib/kubelet/seccomp", "Directory path for seccomp profiles")
	play_kubeCmd.Flags().String("signature-policy", "", "`Pathname` of signature policy file (not usually used)")
	play_kubeCmd.Flags().Bool("start", true, "Start the pod after creating it")
	play_kubeCmd.Flags().Bool("tls-verify", true, "Require HTTPS and verify certificates when contacting registries")
	playCmd.AddCommand(play_kubeCmd)
}
