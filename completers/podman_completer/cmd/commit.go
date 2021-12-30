package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create new image based on the changed container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()
	commitCmd.Flags().StringP("author", "a", "", "Set the author for the image committed")
	commitCmd.Flags().StringArrayP("change", "c", []string{}, "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR")
	commitCmd.Flags().StringP("format", "f", "oci", "`Format` of the image manifest and metadata")
	commitCmd.Flags().String("iidfile", "", "`file` to write the image ID to")
	commitCmd.Flags().Bool("include-volumes", false, "Include container volumes as image volumes")
	commitCmd.Flags().StringP("message", "m", "", "Set commit message for imported image")
	commitCmd.Flags().BoolP("pause", "p", false, "Pause container during commit")
	commitCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	rootCmd.AddCommand(commitCmd)
}
