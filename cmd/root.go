package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcase",
	Short: "Change the 'case' of files and folders",
	Long:  `'gcase' is a CLI tool for renaming files and folders by changing their case`,
}

// Execute runs the root command and exits on error.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
