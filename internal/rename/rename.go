package rename

import (
	"fmt"

	"github.com/shricodev/case/internal/app"
)

func Run(opts app.Options) (app.Result, error) {
	if !opts.Mode.IsValid() {
		return app.Result{}, fmt.Errorf("invalid mode: %q", opts.Mode)
	}

	if !opts.Target.IsValid() {
		return app.Result{}, fmt.Errorf("invalid target: %q", opts.Target)
	}

	// do the thing
	// ...

	return app.Result{}, nil
}
