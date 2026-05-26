package caseconv

import (
	"testing"

	"github.com/shricodev/case/internal/app"
)

func TestApplyMode(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		mode     app.CaseMode
		expected string
	}{
		{"lower", "HELLO", app.ModeLower, "hello"},
		{"upper", "hello", app.ModeUpper, "HELLO"},
		{"capitalize", "hello", app.ModeCapitalize, "Hello"},
		{"capitalize already correct", "Hello", app.ModeCapitalize, "Hello"},
		{"capitalize all upper", "HELLO", app.ModeCapitalize, "HELLO"},
		{"empty string", "", app.ModeLower, ""},
		{"unicode capitalize", "ñoño", app.ModeCapitalize, "Ñoño"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := applyMode(tc.input, tc.mode)
			if got != tc.expected {
				t.Errorf("input=%q mode=%q: got=%q, want=%q", tc.input, tc.mode, got, tc.expected)
			}
		})
	}
}

func TestBuildNewPath(t *testing.T) {
	cases := []struct {
		name        string
		oldPath     string
		mode        app.CaseMode
		kind        app.ItemKind
		preserveExt bool
		expected    string
	}{
		{"lower file", "dir/FILE.TXT", app.ModeLower, app.KindFile, false, "dir/file.txt"},
		{"upper file", "dir/file.txt", app.ModeUpper, app.KindFile, false, "dir/FILE.TXT"},
		{
			"capitalize file",
			"dir/hello.txt",
			app.ModeCapitalize,
			app.KindFile,
			false,
			"dir/Hello.txt",
		},
		{"lower dir", "DIR/SUBDIR", app.ModeLower, app.KindDir, false, "DIR/subdir"},
		{"preserve ext lower", "dir/FILE.TXT", app.ModeLower, app.KindFile, true, "dir/file.TXT"},
		{"preserve ext upper", "dir/file.txt", app.ModeUpper, app.KindFile, true, "dir/FILE.txt"},
		{"preserve ext dir ignored", "DIR/SUBDIR", app.ModeLower, app.KindDir, true, "DIR/subdir"},
		{"no ext preserve ext", "dir/FILE", app.ModeLower, app.KindFile, true, "dir/file"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := BuildNewPath(tc.oldPath, tc.mode, tc.kind, tc.preserveExt)
			if got != tc.expected {
				t.Errorf("oldPath=%q mode=%q preserveExt=%t: got=%q, want=%q",
					tc.oldPath, tc.mode, tc.preserveExt, got, tc.expected)
			}
		})
	}
}
