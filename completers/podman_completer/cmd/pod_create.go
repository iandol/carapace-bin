package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pod_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new empty pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_createCmd).Standalone()
	pod_createCmd.Flags().StringSlice("add-host", []string{}, "Add a custom host-to-IP mapping (host:ip) (default [])")
	pod_createCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	pod_createCmd.Flags().Float64("cpus", 0, "Number of CPUs. The default is 0.000 which means no limit")
	pod_createCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	pod_createCmd.Flags().StringSlice("device", []string{}, "Add a host device to the container")
	pod_createCmd.Flags().StringSlice("device-read-bps", []string{}, "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)")
	pod_createCmd.Flags().StringSlice("dns", []string{}, "Set custom DNS servers")
	pod_createCmd.Flags().StringSlice("dns-opt", []string{}, "Set custom DNS options")
	pod_createCmd.Flags().StringSlice("dns-search", []string{}, "Set custom DNS search domains")
	pod_createCmd.Flags().StringSlice("gidmap", []string{}, "GID map to use for the user namespace")
	pod_createCmd.Flags().StringP("hostname", "h", "", "Set container hostname")
	pod_createCmd.Flags().Bool("infra", true, "Create an infra container associated with the pod to share namespaces with")
	pod_createCmd.Flags().String("infra-command", "", "Overwrite the default ENTRYPOINT of the image")
	pod_createCmd.Flags().String("infra-conmon-pidfile", "", "Path to the file that will receive the PID of conmon")
	pod_createCmd.Flags().String("infra-image", "k8s.gcr.io/pause:3.5", "The image of the infra container to associate with the pod")
	pod_createCmd.Flags().String("infra-name", "", "Assign a name to the container")
	pod_createCmd.Flags().String("ip", "", "Specify a static IPv4 address for the container")
	pod_createCmd.Flags().StringArrayP("label", "l", []string{}, "Set metadata on container")
	pod_createCmd.Flags().StringSlice("label-file", []string{}, "Read in a line delimited file of labels")
	pod_createCmd.Flags().String("mac-address", "", "Container MAC address (e.g. 92:d0:c6:0a:29:33)")
	pod_createCmd.Flags().StringP("name", "n", "", "Assign a name to the pod")
	pod_createCmd.Flags().StringArray("network", []string{}, "Connect a container to a network")
	pod_createCmd.Flags().StringSlice("network-alias", []string{}, "Add network-scoped alias for the container")
	pod_createCmd.Flags().Bool("no-hosts", false, "Do not create /etc/hosts within the container, instead use the version from the image")
	pod_createCmd.Flags().String("pid", "", "PID namespace to use")
	pod_createCmd.Flags().String("pod-id-file", "", "Write the pod ID to the file")
	pod_createCmd.Flags().StringSliceP("publish", "p", []string{}, "Publish a container's port, or a range of ports, to the host (default [])")
	pod_createCmd.Flags().Bool("replace", false, "If a pod with the same name exists, replace it")
	pod_createCmd.Flags().String("share", "cgroup,ipc,net,uts", "A comma delimited list of kernel namespaces the pod will share")
	pod_createCmd.Flags().String("subgidname", "", "Name of range listed in /etc/subgid for use in user namespace")
	pod_createCmd.Flags().String("subuidname", "", "Name of range listed in /etc/subuid for use in user namespace")
	pod_createCmd.Flags().StringSlice("uidmap", []string{}, "UID map to use for the user namespace")
	pod_createCmd.Flags().String("userns", "", "User namespace to use")
	pod_createCmd.Flags().StringArrayP("volume", "v", []string{}, "Bind mount a volume into the container")
	pod_createCmd.Flags().StringArray("volumes-from", []string{}, "Mount volumes from the specified container(s)")
	podCmd.AddCommand(pod_createCmd)
}
