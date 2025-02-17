package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Experimental support for module integration testing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("compact-warnings", false, "Use a more compact representation for warnings")
	testCmd.Flags().String("junit-xml", "", "also write test results to the given file path")
	testCmd.Flags().Bool("no-color", false, "Don't include virtual terminal formatting sequences in the output.")
	rootCmd.AddCommand(testCmd)

	testCmd.Flag("junit-xml").NoOptDefVal = " "

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"junit-xml": carapace.ActionFiles(),
	})
}
