package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shricodev/gcase/internal/app"
)

var upperCmd = &cobra.Command{
	Use:   "upper <path>",
	Short: "Rename items to uppercase",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCaseCommand(cmd, app.ModeUpper, args, opts)
	},
}

func init() {
	addRenameFlags(upperCmd, &opts)
	rootCmd.AddCommand(upperCmd)
}
