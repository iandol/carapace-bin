package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()
	loginCmd.Flags().String("authfile", "", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	loginCmd.Flags().String("cert-dir", "", "use certificates at the specified path to access the registry")
	loginCmd.Flags().Bool("get-login", false, "Return the current login user for the registry")
	loginCmd.Flags().StringP("password", "p", "", "Password for registry")
	loginCmd.Flags().Bool("password-stdin", false, "Take the password from stdin")
	loginCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	loginCmd.Flags().StringP("username", "u", "", "Username for registry")
	loginCmd.Flags().BoolP("verbose", "v", false, "Write more detailed information to stdout")
	rootCmd.AddCommand(loginCmd)
}
