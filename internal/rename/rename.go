package rename

import (
	"fmt"
	"os"

	"github.com/shricodev/case/internal/app"
	"github.com/shricodev/case/internal/walker"
)

func Run(opts app.Options) (app.Result, error) {
	if !opts.Mode.IsValid() {
		return app.Result{}, fmt.Errorf("invalid mode: %q", opts.Mode)
	}

	if !opts.Target.IsValid() {
		return app.Result{}, fmt.Errorf("invalid target: %q", opts.Target)
	}

	items, err := walker.Collect(walker.Options{
		Root:          opts.Root,
		Target:        opts.Target,
		Recursive:     opts.Recursive,
		IncludeHidden: opts.IncludeHidden,
	})
	if err != nil {
		return app.Result{}, fmt.Errorf("error collecting paths: %w", err)
	}

	for _, item := range items {
		fmt.Fprintf(os.Stdout, "%s %s\n", item.Kind, item.Path)
	}

	return app.Result{}, nil
}
