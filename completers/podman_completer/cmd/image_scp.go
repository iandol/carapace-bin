package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_scpCmd = &cobra.Command{
	Use:   "scp",
	Short: "securely copy images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_scpCmd).Standalone()
	image_scpCmd.Flags().BoolP("quiet", "q", false, "Suppress the output")
	imageCmd.AddCommand(image_scpCmd)
}
