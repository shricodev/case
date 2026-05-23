package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shricodev/case/internal/app"
	"github.com/shricodev/case/internal/rename"
)

func runCaseCommand(cmd *cobra.Command, mode app.CaseMode, args []string) error {
	// since the arg count is already validated in the command, this is safe.
	// But still for the sake of safety
	if len(args) == 0 {
		return fmt.Errorf("missing path")
	}

	opts.Root = args[0]
	opts.Mode = mode

	result, err := rename.Run(opts)
	if err != nil {
		return err
	}

	return printResult(cmd.OutOrStdout(), result)
}
