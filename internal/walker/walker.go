package walker

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/shricodev/gcase/internal/app"
)

// Collect gathers filesystem entries under opts.Root that match opts.
func Collect(opts Options) ([]Item, error) {
	opts.Root = filepath.Clean(opts.Root)
	return walk(opts)
}

func walk(opts Options) ([]Item, error) {
	// there's no way to figure out all the paths that's going to come in
	// so we will use the good old append() without assigning a capacity.
	// increases the TC but that's fine for now.
	items := []Item{}

	root := filepath.Clean(opts.Root)

	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	addItem := func(path string, kind app.ItemKind) {
		if !opts.IncludeHidden && isHiddenUnix(path) {
			return
		}

		if !matchTarget(kind, opts.Target) {
			return
		}

		items = append(items, Item{
			Path: path,
			Kind: kind,
		})
	}

	if !info.IsDir() {
		addItem(root, app.KindFile)
		return items, nil
	}

	if opts.Recursive {
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}

			// Do not collect "." itself.
			// Example: gcase lower . -r should not try to rename "."
			if path == root && isCurrentOrParentDir(path) {
				return nil
			}

			// If hidden dirs are disabled, skip the whole hidden directory.
			// This avoids walking into things like .git, .cache, .next, etc.
			if !opts.IncludeHidden && d.IsDir() && isHiddenUnix(path) {
				return filepath.SkipDir
			}

			kind := kindFromDirEntry(d)
			addItem(path, kind)

			return nil
		})
		if err != nil {
			return nil, err
		}

		return items, nil
	}

	// Non-recursive mode.
	// First, include the root directory itself, unless it is "." or "..".
	if !isCurrentOrParentDir(root) {
		addItem(root, app.KindDir)
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		path := filepath.Join(root, entry.Name())
		kind := kindFromDirEntry(entry)

		addItem(path, kind)
	}

	return items, nil
}

func kindFromDirEntry(entry os.DirEntry) app.ItemKind {
	if entry.IsDir() {
		return app.KindDir
	}
	return app.KindFile
}

func isHiddenUnix(path string) bool {
	name := filepath.Base(path)
	if name == "." || name == ".." {
		return false
	}

	return strings.HasPrefix(name, ".")
}

func isCurrentOrParentDir(path string) bool {
	clean := filepath.Clean(path)
	return clean == "." || clean == ".."
}

func matchTarget(kind app.ItemKind, target app.Target) bool {
	switch target {
	case app.TargetAll:
		return true
	case app.TargetFiles:
		return kind == app.KindFile
	case app.TargetDirs:
		return kind == app.KindDir
	default:
		return false
	}
}
