package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secret_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect a secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_inspectCmd).Standalone()
	secret_inspectCmd.Flags().String("format", "", "Format volume output using Go template")
	secretCmd.AddCommand(secret_inspectCmd)
}
