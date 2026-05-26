package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shricodev/gcase/internal/app"
)

// capitalizeCmd represents the capitalize command
var capitalizeCmd = &cobra.Command{
	Use:   "capitalize <path>",
	Short: "Capitalize item names",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCaseCommand(cmd, app.ModeCapitalize, args, opts)
	},
}

func init() {
	addRenameFlags(capitalizeCmd, &opts)
	rootCmd.AddCommand(capitalizeCmd)
}
