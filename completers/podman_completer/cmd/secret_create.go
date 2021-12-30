package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secret_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_createCmd).Standalone()
	secret_createCmd.Flags().String("driver", "file", "Specify secret driver")
	secret_createCmd.Flags().String("driver-opts", "", "Specify driver specific options")
	secret_createCmd.Flags().Bool("env", false, "Read secret data from environment variable")
	secretCmd.AddCommand(secret_createCmd)
}
