package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_trust_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display trust policy for the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_trust_showCmd).Standalone()
	image_trust_showCmd.Flags().BoolP("json", "j", false, "Output as json")
	image_trust_showCmd.Flags().String("policypath", "", "")
	image_trust_showCmd.Flags().Bool("raw", false, "Output raw policy file")
	image_trust_showCmd.Flags().String("registrypath", "", "")
	image_trustCmd.AddCommand(image_trust_showCmd)
}
