package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_createCmd = &cobra.Command{
	Use:   "create",
	Short: "network create",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_createCmd).Standalone()
	network_createCmd.Flags().Bool("disable-dns", false, "disable dns plugin")
	network_createCmd.Flags().StringP("driver", "d", "bridge", "driver to manage the network")
	network_createCmd.Flags().String("gateway", "", "IPv4 or IPv6 gateway for the subnet")
	network_createCmd.Flags().Bool("internal", false, "restrict external access from this network")
	network_createCmd.Flags().String("ip-range", "", "allocate container IP from range")
	network_createCmd.Flags().Bool("ipv6", false, "enable IPv6 networking")
	network_createCmd.Flags().StringArray("label", []string{}, "set metadata on a network")
	network_createCmd.Flags().String("macvlan", "", "create a Macvlan connection based on this device")
	network_createCmd.Flags().StringArrayP("opt", "o", []string{}, "Set driver specific options (default [])")
	network_createCmd.Flags().String("subnet", "", "subnet in CIDR format")
	networkCmd.AddCommand(network_createCmd)
}
