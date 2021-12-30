package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secret_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_rmCmd).Standalone()
	secret_rmCmd.Flags().BoolP("all", "a", false, "Remove all secrets")
	secretCmd.AddCommand(secret_rmCmd)
}
