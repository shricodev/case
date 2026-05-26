package rename

import (
	"testing"

	"github.com/shricodev/case/internal/app"
)

func TestSortByDepth(t *testing.T) {
	items := []app.ItemResult{
		{OldPath: "a/b/c/file.txt"},
		{OldPath: "a/file.txt"},
		{OldPath: "a/b/file.txt"},
	}

	sortByDepth(items)

	expected := []string{
		"a/b/c/file.txt",
		"a/b/file.txt",
		"a/file.txt",
	}

	for i, item := range items {
		if item.OldPath != expected[i] {
			t.Errorf("index=%d: got=%s, want=%q", i, item.OldPath, expected[i])
		}
	}
}
