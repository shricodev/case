package walker

import "github.com/shricodev/gcase/internal/app"

// Item is a single filesystem entry collected during a walk.
type Item struct {
	Path string
	Kind app.ItemKind
}

// Options configures a directory walk.
type Options struct {
	Root          string
	Target        app.Target
	Recursive     bool
	IncludeHidden bool
}
