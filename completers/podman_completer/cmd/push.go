package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an image to a specified destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()
	pushCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	pushCmd.Flags().String("cert-dir", "", "Path to a directory containing TLS certificates and keys")
	pushCmd.Flags().Bool("compress", false, "Compress tarball image layers when pushing to a directory using the 'dir' transport. (default is same compression type as source)")
	pushCmd.Flags().String("compression-format", "", "compression format to use")
	pushCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	pushCmd.Flags().String("digestfile", "", "Write the digest of the pushed image to the specified file")
	pushCmd.Flags().Bool("disable-content-trust", false, "This is a Docker specific option and is a NOOP")
	pushCmd.Flags().StringP("format", "f", "", "Manifest type (oci, v2s2, or v2s1) to use in the destination (default is manifest type of source, with fallbacks)")
	pushCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pushing images")
	pushCmd.Flags().Bool("remove-signatures", false, "Discard any pre-existing signatures in the image")
	pushCmd.Flags().String("sign-by", "", "Add a signature at the destination using the specified key")
	pushCmd.Flags().String("signature-policy", "", "Path to a signature-policy file")
	pushCmd.Flags().Bool("tls-verify", true, "Require HTTPS and verify certificates when contacting registries")
	rootCmd.AddCommand(pushCmd)
}
