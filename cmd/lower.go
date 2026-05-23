package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shricodev/case/internal/app"
)

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower <path>",
	Short: "Rename items to lowercase",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCaseCommand(cmd, app.ModeLower, args, opts)
	},
}

func init() {
	addRenameFlags(lowerCmd, &opts)
	rootCmd.AddCommand(lowerCmd)
}
