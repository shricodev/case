package rename

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/shricodev/case/internal/app"
	"github.com/shricodev/case/internal/caseconv"
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

	for i := range result.Items {
		item := &result.Items[i]
		newPath := caseconv.BuildNewPath(item.OldPath, opts.Mode, item.Kind, opts.PreserveExt)
		item.NewPath = newPath

		if item.OldPath == newPath {
			item.Status = app.StatusSkipped
			continue
		}

		if opts.DryRun {
			item.Status = app.StatusPlanned
		} else {
			if err := os.Rename(item.OldPath, item.NewPath); err != nil {
				item.Error = err
				item.Status = app.StatusFailed
			} else {
				item.Status = app.StatusRenamed
			}
		}
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
	// also since we're doing it all locally, and not modifying the real path, this is fine.
	// and still cross-compatible to all OS even windows.
	path = strings.ReplaceAll(path, `\`, `/`)
	path = strings.Trim(path, `/`)

	if path == "" {
		return 0
	}

	return strings.Count(path, `/`)
}
