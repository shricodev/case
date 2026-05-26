package walker

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/shricodev/gcase/internal/app"
)

func TestIsHiddenUnix(t *testing.T) {
	t.Run("Hidden", func(t *testing.T) {
		cases := []struct {
			path     string
			expected bool
		}{
			{".git", true},
			{".env", true},
			{"/some/path/.hidden", true},
		}

		for _, tc := range cases {
			t.Run(tc.path, func(t *testing.T) {
				got := isHiddenUnix(tc.path)
				if got != tc.expected {
					t.Errorf("got=%t, want=%t", got, tc.expected)
				}
			})
		}
	})

	t.Run("Not Hidden", func(t *testing.T) {
		cases := []struct {
			path     string
			expected bool
		}{
			{"foo.txt", false},
			{".", false},
			{"..", false},
			{"foo", false},
			{"/some/path/visible", false},
		}

		for _, tc := range cases {
			t.Run(tc.path, func(t *testing.T) {
				got := isHiddenUnix(tc.path)
				if got != tc.expected {
					t.Errorf("got=%t, want=%t", got, tc.expected)
				}
			})
		}
	})
}

func TestIsCurrentOrParentDir(t *testing.T) {
	cases := []struct {
		path     string
		expected bool
	}{
		{".", true},
		{"..", true},
		{"foo.txt", false},
		{"foo/something", false},
	}

	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			got := isCurrentOrParentDir(tc.path)
			if got != tc.expected {
				t.Errorf("got=%t, want=%t", got, tc.expected)
			}
		})
	}
}

func assertPaths(t *testing.T, items []Item, expected []string) {
	t.Helper()

	got := make([]string, len(items))
	for i, item := range items {
		got[i] = item.Path
	}

	sort.Strings(got)
	sort.Strings(expected)

	if len(got) != len(expected) {
		t.Errorf(
			"count: got %d, want %d\ngot:  %v\nwant: %v",
			len(got),
			len(expected),
			got,
			expected,
		)
		return
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("path[%d]: got %q, want %q", i, got[i], expected[i])
		}
	}
}

func makeFile(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatal(err)
	}
}

func makeDir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
}

func TestMatchTarget(t *testing.T) {
	cases := []struct {
		name     string
		kind     app.ItemKind
		target   app.Target
		expected bool
	}{
		{"file matches TargetAll", app.KindFile, app.TargetAll, true},
		{"dir matches TargetAll", app.KindDir, app.TargetAll, true},
		{"file matches TargetFiles", app.KindFile, app.TargetFiles, true},
		{"dir does not match TargetFiles", app.KindDir, app.TargetFiles, false},
		{"dir matches TargetDirs", app.KindDir, app.TargetDirs, true},
		{"file does not match TargetDirs", app.KindFile, app.TargetDirs, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := matchTarget(tc.kind, tc.target)
			if got != tc.expected {
				t.Errorf(
					"kind=%q, target=%q, got=%t, want=%t",
					tc.kind,
					tc.target,
					got,
					tc.expected,
				)
			}
		})
	}
}

func TestCollect_SingleFile(t *testing.T) {
	root := t.TempDir()
	f := filepath.Join(root, "only.txt")
	makeFile(t, f)

	items, err := Collect(Options{Root: f, Target: app.TargetAll})
	if err != nil {
		t.Fatal(err)
	}

	assertPaths(t, items, []string{f})

	if items[0].Kind != app.KindFile {
		t.Errorf("kind: got %q, want %q", items[0].Kind, app.KindFile)
	}
}

func TestCollect_NonRecursive_DotRoot(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeFile(t, filepath.Join(root, "b.txt"))
	makeDir(t, filepath.Join(root, "subdir"))

	t.Chdir(root)

	items, err := Collect(Options{Root: ".", Target: app.TargetAll})
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range items {
		if item.Path == "." {
			t.Error("root '.' must not appear in results")
		}
	}

	assertPaths(t, items, []string{"a.txt", "b.txt", "subdir"})
}

func TestCollect_NonRecursive_NamedRoot(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeFile(t, filepath.Join(root, "b.txt"))
	makeDir(t, filepath.Join(root, "subdir"))

	items, err := Collect(Options{Root: root, Target: app.TargetAll})
	if err != nil {
		t.Fatal(err)
	}

	assertPaths(t, items, []string{
		root,
		filepath.Join(root, "a.txt"),
		filepath.Join(root, "b.txt"),
		filepath.Join(root, "subdir"),
	})
}

func TestCollect_NonRecursive_TargetFiles(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeFile(t, filepath.Join(root, "b.txt"))
	makeDir(t, filepath.Join(root, "subdir"))

	items, err := Collect(Options{Root: root, Target: app.TargetFiles})
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range items {
		if item.Kind != app.KindFile {
			t.Errorf("expected only files, got kind=%q for path=%q", item.Kind, item.Path)
		}
	}

	assertPaths(t, items, []string{
		filepath.Join(root, "a.txt"),
		filepath.Join(root, "b.txt"),
	})
}

func TestCollect_NonRecursive_TargetDirs(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeDir(t, filepath.Join(root, "subdir"))
	makeDir(t, filepath.Join(root, "another"))

	items, err := Collect(Options{Root: root, Target: app.TargetDirs})
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range items {
		if item.Kind != app.KindDir {
			t.Errorf("expected only dirs, got kind=%q for path=%q", item.Kind, item.Path)
		}
	}

	assertPaths(t, items, []string{
		root,
		filepath.Join(root, "subdir"),
		filepath.Join(root, "another"),
	})
}

func TestCollect_Recursive_NoHidden(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeFile(t, filepath.Join(root, "subdir", "b.txt"))
	makeFile(t, filepath.Join(root, ".hidden_file"))
	makeFile(t, filepath.Join(root, ".hidden_dir", "child.txt"))

	items, err := Collect(Options{
		Root:          root,
		Target:        app.TargetAll,
		Recursive:     true,
		IncludeHidden: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertPaths(t, items, []string{
		root,
		filepath.Join(root, "a.txt"),
		filepath.Join(root, "subdir"),
		filepath.Join(root, "subdir", "b.txt"),
	})
}

func TestCollect_Recursive_IncludeHidden(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeFile(t, filepath.Join(root, ".hidden_file"))
	makeFile(t, filepath.Join(root, ".hidden_dir", "child.txt"))

	items, err := Collect(Options{
		Root:          root,
		Target:        app.TargetAll,
		Recursive:     true,
		IncludeHidden: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertPaths(t, items, []string{
		root,
		filepath.Join(root, "a.txt"),
		filepath.Join(root, ".hidden_file"),
		filepath.Join(root, ".hidden_dir"),
		filepath.Join(root, ".hidden_dir", "child.txt"),
	})
}

func TestCollect_Recursive_TargetFiles(t *testing.T) {
	root := t.TempDir()
	makeFile(t, filepath.Join(root, "a.txt"))
	makeFile(t, filepath.Join(root, "subdir", "b.txt"))
	makeDir(t, filepath.Join(root, "emptydir"))

	items, err := Collect(Options{
		Root:      root,
		Target:    app.TargetFiles,
		Recursive: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range items {
		if item.Kind != app.KindFile {
			t.Errorf("expected only files, got kind=%q for path=%q", item.Kind, item.Path)
		}
	}

	assertPaths(t, items, []string{
		filepath.Join(root, "a.txt"),
		filepath.Join(root, "subdir", "b.txt"),
	})
}

func TestCollect_NonExistentRoot(t *testing.T) {
	_, err := Collect(Options{Root: "/nonexistent/path/xyz", Target: app.TargetAll})
	if err == nil {
		t.Error("expected an error for a non-existent root, got nil")
	}
}
