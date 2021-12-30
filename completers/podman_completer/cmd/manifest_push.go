package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a manifest list or image index to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_pushCmd).Standalone()
	manifest_pushCmd.Flags().Bool("all", true, "also push the images in the list")
	manifest_pushCmd.Flags().String("authfile", "", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	manifest_pushCmd.Flags().String("cert-dir", "", "use certificates at the specified path to access the registry")
	manifest_pushCmd.Flags().String("creds", "", "use `[username[:password]]` for accessing the registry")
	manifest_pushCmd.Flags().String("digestfile", "", "after copying the image, write the digest of the resulting digest to the file")
	manifest_pushCmd.Flags().StringP("format", "f", "", "manifest type (oci or v2s2) to attempt to use when pushing the manifest list (default is manifest type of source)")
	manifest_pushCmd.Flags().BoolP("quiet", "q", false, "don't output progress information when pushing lists")
	manifest_pushCmd.Flags().Bool("remove-signatures", false, "don't copy signatures when pushing images")
	manifest_pushCmd.Flags().Bool("rm", false, "remove the manifest list if push succeeds")
	manifest_pushCmd.Flags().String("sign-by", "", "sign the image using a GPG key with the specified `FINGERPRINT`")
	manifest_pushCmd.Flags().Bool("tls-verify", true, "require HTTPS and verify certificates when accessing the registry")
	manifestCmd.AddCommand(manifest_pushCmd)
}
