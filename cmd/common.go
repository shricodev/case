package cmd

import (
	"github.com/spf13/cobra"

	"github.com/shricodev/gcase/internal/app"
	"github.com/shricodev/gcase/internal/rename"
)

// runCaseCommand wires a cobra command to rename.Run for the given mode.
func runCaseCommand(cmd *cobra.Command, mode app.CaseMode, args []string, opts app.Options) error {
	opts.Root = args[0]
	opts.Mode = mode

	result, err := rename.Run(opts)
	if err != nil {
		return err
	}

	// fmt.Printf("sorted result: %+v\n", result)
	return printResult(cmd.OutOrStdout(), result)
}
