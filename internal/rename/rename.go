package rename

import (
	"fmt"
	"slices"
	"strings"

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

	itemResults := make([]app.ItemResult, len(items))
	for i, item := range items {
		itemResults[i] = app.ItemResult{
			OldPath: item.Path,
			Kind:    item.Kind,
		}
	}

	sortByDepth(itemResults)
	result := app.Result{
		Items:  itemResults,
		DryRun: opts.DryRun,
	}

	return result, nil
}

func sortByDepth(items []app.ItemResult) {
	slices.SortFunc(items, func(x, y app.ItemResult) int {
		xLen := pathDepth(x.OldPath)
		yLen := pathDepth(y.OldPath)

		return yLen - xLen
	})
}

func pathDepth(path string) int {
	// NOTE: we dont want to use filepath.Separator, as what if the user is in
	// a windows machine, working in a linux style file system or vice versa?
	// so this should be safer.
	path = strings.ReplaceAll(path, `\`, `/`)
	path = strings.Trim(path, `/`)

	if path == "" {
		return 0
	}

	return strings.Count(path, `/`)
}
