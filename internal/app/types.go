package app

// CaseMode specifies how entry names are transformed.
type CaseMode string

const (
	ModeLower      CaseMode = "lower"      // all characters to lowercase
	ModeUpper      CaseMode = "upper"      // all characters to uppercase
	ModeCapitalize CaseMode = "capitalize" // first character to uppercase
)

// IsValid reports whether m is a recognized CaseMode.
func (m CaseMode) IsValid() bool {
	switch m {
	case ModeLower, ModeUpper, ModeCapitalize:
		return true
	default:
		return false
	}
}

// Target controls which filesystem entry types are renamed.
type Target string

const (
	TargetDirs  Target = "dirs"  // directories only
	TargetFiles Target = "files" // files only
	TargetAll   Target = "all"   // both files and directories
)

// IsValid reports whether t is a recognized Target.
func (t Target) IsValid() bool {
	switch t {
	case TargetAll, TargetDirs, TargetFiles:
		return true
	default:
		return false
	}
}

// Options holds the configuration for a rename run.
type Options struct {
	Root          string
	Mode          CaseMode
	Target        Target
	Recursive     bool
	DryRun        bool
	IncludeHidden bool
	PreserveExt   bool // when true, file extensions are left unchanged
}

// ItemStatus describes the outcome of a single rename attempt.
type ItemStatus string

const (
	StatusPlanned ItemStatus = "planned" // dry-run: would be renamed
	StatusFailed  ItemStatus = "failed"  // rename failed
	StatusRenamed ItemStatus = "renamed" // successfully renamed
	StatusSkipped ItemStatus = "skipped" // old and new names are identical
)

// ItemKind distinguishes files from directories.
type ItemKind string

const (
	KindFile ItemKind = "file"
	KindDir  ItemKind = "dir"
)

// ItemResult holds the before/after paths and outcome for one entry.
type ItemResult struct {
	OldPath string
	NewPath string
	Kind    ItemKind
	Status  ItemStatus
	Error   error
}

// Result holds the full output of a rename run.
type Result struct {
	Items  []ItemResult
	DryRun bool
}

// Count returns the number of items with the given status.
func (r Result) Count(status ItemStatus) int {
	count := 0
	for _, item := range r.Items {
		if item.Status == status {
			count++
		}
	}

	return count
}
