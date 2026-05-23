package app

type CaseMode string

const (
	ModeLower      CaseMode = "lower"
	ModeUpper      CaseMode = "upper"
	ModeCapitalize CaseMode = "capitalize"
)

func (m CaseMode) IsValid() bool {
	switch m {
	case ModeLower, ModeUpper, ModeCapitalize:
		return true
	default:
		return false
	}
}

type Target string

const (
	TargetDirs  Target = "dirs"
	TargetFiles Target = "files"
	TargetAll   Target = "all"
)

func (t Target) IsValid() bool {
	switch t {
	case TargetAll, TargetDirs, TargetFiles:
		return true
	default:
		return false
	}
}

type Options struct {
	Root          string
	Mode          CaseMode
	Target        Target
	Recursive     bool
	DryRun        bool
	IncludeHidden bool
	PreserveExt   bool
}

type ItemStatus string

const (
	StatusPlanned ItemStatus = "planned"
	StatusFailed  ItemStatus = "failed"
	StatusRenamed ItemStatus = "renamed"
	StatusSkipped ItemStatus = "skipped"
)

type ItemKind string

const (
	KindFile ItemKind = "file"
	KindDir  ItemKind = "dir"
)

type ItemResult struct {
	OldPath string
	NewPath string
	Kind    ItemKind
	Status  ItemStatus
	Error   error
}

type Result struct {
	Items  []ItemResult
	DryRun bool
}

func (r Result) Count(status ItemStatus) int {
	count := 0
	for _, item := range r.Items {
		if item.Status == status {
			count++
		}
	}

	return count
}
