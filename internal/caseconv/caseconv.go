package caseconv

import (
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/shricodev/gcase/internal/app"
)

func applyMode(name string, mode app.CaseMode) string {
	switch mode {
	case app.ModeLower:
		return strings.ToLower(name)
	case app.ModeUpper:
		return strings.ToUpper(name)
	case app.ModeCapitalize:
		// strings.Title is deprecated so this is the way for now.
		if name == "" {
			return ""
		}

		r, size := utf8.DecodeRuneInString(name)
		return string(unicode.ToUpper(r)) + name[size:]
	default:
		return ""
	}
}

// BuildNewPath returns the renamed path for oldPath under mode.
// When preserveExt is true and kind is KindFile, the extension is left unchanged.
func BuildNewPath(
	oldPath string,
	mode app.CaseMode,
	kind app.ItemKind,
	preserveExt bool,
) string {
	dir := filepath.Dir(oldPath)
	base := filepath.Base(oldPath)

	// this is a file
	if preserveExt && kind == app.KindFile {
		ext := filepath.Ext(base)
		stem := strings.TrimSuffix(base, ext)
		return filepath.Join(dir, applyMode(stem, mode)+ext)
	}

	return filepath.Join(dir, applyMode(base, mode))
}
