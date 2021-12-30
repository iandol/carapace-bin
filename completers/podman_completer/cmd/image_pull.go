package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pullCmd).Standalone()
	image_pullCmd.Flags().Bool("all-tags", false, "All tagged images in the repository will be pulled")
	image_pullCmd.Flags().String("arch", "", "Use `ARCH` instead of the architecture of the machine for choosing images")
	image_pullCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	image_pullCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	image_pullCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	image_pullCmd.Flags().Bool("disable-content-trust", false, "This is a Docker specific option and is a NOOP")
	image_pullCmd.Flags().String("os", "", "Use `OS` instead of the running OS for choosing images")
	image_pullCmd.Flags().String("platform", "", "Specify the platform for selecting the image.  (Conflicts with arch and os)")
	image_pullCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pulling images")
	image_pullCmd.Flags().String("signature-policy", "", "`Pathname` of signature policy file (not usually used)")
	image_pullCmd.Flags().Bool("tls-verify", true, "Require HTTPS and verify certificates when contacting registries")
	image_pullCmd.Flags().String("variant", "", "Use VARIANT instead of the running architecture variant for choosing images")
	imageCmd.AddCommand(image_pullCmd)
}
