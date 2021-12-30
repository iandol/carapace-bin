package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()
	pullCmd.Flags().Bool("all-tags", false, "All tagged images in the repository will be pulled")
	pullCmd.Flags().String("arch", "", "Use `ARCH` instead of the architecture of the machine for choosing images")
	pullCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	pullCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	pullCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	pullCmd.Flags().Bool("disable-content-trust", false, "This is a Docker specific option and is a NOOP")
	pullCmd.Flags().String("os", "", "Use `OS` instead of the running OS for choosing images")
	pullCmd.Flags().String("platform", "", "Specify the platform for selecting the image.  (Conflicts with arch and os)")
	pullCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pulling images")
	pullCmd.Flags().String("signature-policy", "", "`Pathname` of signature policy file (not usually used)")
	pullCmd.Flags().Bool("tls-verify", true, "Require HTTPS and verify certificates when contacting registries")
	pullCmd.Flags().String("variant", "", "Use VARIANT instead of the running architecture variant for choosing images")
	rootCmd.AddCommand(pullCmd)
}
