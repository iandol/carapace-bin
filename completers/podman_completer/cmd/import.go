package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a tarball to create a filesystem image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()
	importCmd.Flags().StringArrayP("change", "c", []string{}, "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR")
	importCmd.Flags().StringP("message", "m", "", "Set commit message for imported image")
	importCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	importCmd.Flags().String("signature-policy", "", "Path to a signature-policy file")
	rootCmd.AddCommand(importCmd)
}
