package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/shricodev/case/internal/app"
)

var opts app.Options

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "case",
	Short: "Change the 'case' of files and folders",
	Long:  `'case' is a CLI tool for renaming files and folders by changing their case`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
