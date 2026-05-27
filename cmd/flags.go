package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shricodev/gcase/internal/app"
)

var opts app.Options

// addRenameFlags registers the shared rename flags on cmd, storing values in opts.
func addRenameFlags(cmd *cobra.Command, opts *app.Options) {
	cmd.Flags().
		BoolVarP(&opts.DryRun, "dry-run", "n", false, "show what would be renamed without changing anything")
	cmd.Flags().
		BoolVar(&opts.IncludeHidden, "include-hidden", false, "include hidden files and directories")
	cmd.Flags().
		BoolVarP(&opts.Recursive, "recursive", "r", false, "rename recursively")
	cmd.Flags().
		BoolVar(&opts.PreserveExt, "preserve-extension", true, "preserve file extensions when renaming files")
	cmd.Flags().
		StringVar((*string)(&opts.Target), "target", "dirs", "what to rename: dirs, files, or all")
}
