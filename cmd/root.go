package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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

func init() {
	rootCmd.PersistentFlags().
		BoolVarP(&dryRun, "dry-run", "n", false, "show what would be renamed without changing anything")
	rootCmd.PersistentFlags().
		BoolVar(&includeHidden, "include-hidden", false, "include hidden files and directories")
	rootCmd.PersistentFlags().
		BoolVarP(&recursive, "recursive", "r", false, "rename recursively")
	rootCmd.PersistentFlags().
		BoolVar(&preserveExt, "preserve-extension", true, "preserve file extensions when renaming files")
	rootCmd.PersistentFlags().
		StringVar(&target, "target", "dirs", "what to renamme: dirs, files, or all")
}
