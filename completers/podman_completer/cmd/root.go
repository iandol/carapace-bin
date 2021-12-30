package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "podman",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().String("cgroup-manager", "cgroupfs", "Cgroup manager to use (\"cgroupfs\"|\"systemd\")")
	rootCmd.PersistentFlags().String("cni-config-dir", "", "Path of the configuration directory for CNI networks")
	rootCmd.PersistentFlags().String("conmon", "", "Path of the conmon binary")
	rootCmd.Flags().StringP("connection", "c", "", "Connection to use for remote Podman service")
	rootCmd.Flags().String("context", "default", "Name of the context to use to connect to the daemon (This flag is a NOOP and provided solely for scripting compatibility.)")
	rootCmd.PersistentFlags().String("cpu-profile", "", "Path for the cpu-profiling results")
	rootCmd.PersistentFlags().String("default-mounts-file", "", "Path to default mounts file")
	rootCmd.PersistentFlags().String("events-backend", "file", "Events backend to use (\"file\"|\"journald\"|\"none\")")
	rootCmd.PersistentFlags().Bool("help", false, "Help for podman")
	rootCmd.PersistentFlags().StringSlice("hooks-dir", []string{"/usr/share/containers/oci/hooks.d"}, "Set the OCI hooks directory path (may be set multiple times)")
	rootCmd.Flags().String("identity", "", "path to SSH identity file, (CONTAINER_SSHKEY)")
	rootCmd.PersistentFlags().String("log-level", "warn", "Log messages above specified level (trace, debug, info, warn, warning, error, fatal, panic)")
	rootCmd.PersistentFlags().Int("max-workers", 13, "The maximum number of workers for parallel operations")
	rootCmd.PersistentFlags().String("memory-profile", "", "Path for the memory-profiling results")
	rootCmd.PersistentFlags().String("namespace", "", "Set the libpod namespace, used to create separate views of the containers and pods on the system")
	rootCmd.PersistentFlags().String("network-backend", "cni", "Network backend to use (\"cni\"|\"netavark\")")
	rootCmd.PersistentFlags().String("network-cmd-path", "", "Path to the command for configuring the network")
	rootCmd.PersistentFlags().String("registries-conf", "", "Path to a registries.conf to use for image processing")
	rootCmd.Flags().BoolP("remote", "r", false, "Access remote Podman service")
	rootCmd.PersistentFlags().String("root", "", "Path to the root directory in which data, including images, is stored")
	rootCmd.PersistentFlags().String("runroot", "", "Path to the 'run directory' where all state information is stored")
	rootCmd.PersistentFlags().String("runtime", "", "Path to the OCI-compatible binary used to run containers, default is /usr/bin/runc")
	rootCmd.PersistentFlags().StringArray("runtime-flag", []string{}, "add global flags for the container runtime")
	rootCmd.PersistentFlags().String("storage-driver", "", "Select which storage driver is used to manage storage of images and containers (default is overlay)")
	rootCmd.PersistentFlags().StringArray("storage-opt", []string{}, "Used to pass an option to the storage driver")
	rootCmd.PersistentFlags().Bool("syslog", false, "Output logging information to syslog as well as the console (default false)")
	rootCmd.PersistentFlags().String("tmpdir", "", "Path to the tmp directory for libpod state content.")
	rootCmd.PersistentFlags().Bool("trace", false, "Enable opentracing output (default false)")
	rootCmd.Flags().String("url", "unix:/run/user/1000/podman/podman.sock", "URL to access Podman service (CONTAINER_HOST)")
	rootCmd.Flags().BoolP("version", "v", false, "version for podman")
}
