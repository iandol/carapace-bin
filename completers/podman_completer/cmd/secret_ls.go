package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secret_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_lsCmd).Standalone()
	secret_lsCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter secret output")
	secret_lsCmd.Flags().String("format", "{{.ID}}	{{.Name}}	{{.Driver}}	{{.CreatedAt}}	{{.UpdatedAt}}", "Format volume output using Go template")
	secret_lsCmd.Flags().Bool("noheading", false, "Do not print headers")
	secretCmd.AddCommand(secret_lsCmd)
}
