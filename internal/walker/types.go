package walker

import "github.com/shricodev/case/internal/app"

type Item struct {
	Path string
	Kind app.ItemKind
}

type Options struct {
	Root          string
	Target        app.Target
	Recursive     bool
	IncludeHidden bool
}
