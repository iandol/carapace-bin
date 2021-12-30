package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var manifest_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add images to a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_addCmd).Standalone()
	manifest_addCmd.Flags().Bool("all", false, "add all of the list's images if the image is a list")
	manifest_addCmd.Flags().StringSlice("annotation", []string{}, "set an `annotation` for the specified image")
	manifest_addCmd.Flags().String("arch", "", "override the `architecture` of the specified image")
	manifest_addCmd.Flags().String("authfile", "", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	manifest_addCmd.Flags().String("cert-dir", "", "use certificates at the specified path to access the registry")
	manifest_addCmd.Flags().String("creds", "", "use `[username[:password]]` for accessing the registry")
	manifest_addCmd.Flags().StringSlice("features", []string{}, "override the `features` of the specified image")
	manifest_addCmd.Flags().String("os", "", "override the `OS` of the specified image")
	manifest_addCmd.Flags().String("os-version", "", "override the OS `version` of the specified image")
	manifest_addCmd.Flags().Bool("tls-verify", true, "require HTTPS and verify certificates when accessing the registry")
	manifest_addCmd.Flags().String("variant", "", "override the `Variant` of the specified image")
	manifestCmd.AddCommand(manifest_addCmd)
}
