package cmd

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/shricodev/gcase/internal/app"
)

// messageForStatus returns a short human-readable label for status.
func messageForStatus(status app.ItemStatus) string {
	switch status {
	case app.StatusPlanned:
		return "would rename"
	case app.StatusFailed:
		return "failed"
	case app.StatusRenamed:
		return "renamed"
	case app.StatusSkipped:
		return "skipped"
	default:
		return "unknown"
	}
}

// printResult writes a formatted, aligned summary of result to w.
func printResult(w io.Writer, result app.Result) error {
	// find the longest label and old base name for alignment
	maxLabel := 0
	maxOldBase := 0
	for _, item := range result.Items {
		label := messageForStatus(item.Status)
		if len(label) > maxLabel {
			maxLabel = len(label)
		}
		oldBase := filepath.Base(item.OldPath)
		if len(oldBase) > maxOldBase {
			maxOldBase = len(oldBase)
		}
	}

	for _, item := range result.Items {
		label := messageForStatus(item.Status)
		oldBase := filepath.Base(item.OldPath)

		if item.Status == app.StatusFailed && item.Error != nil {
			_, err := fmt.Fprintf(w, "%-*s  %s: %v\n", maxLabel, label, oldBase, item.Error)
			if err != nil {
				return err
			}
			continue
		}

		newBase := filepath.Base(item.NewPath)
		_, err := fmt.Fprintf(w, "%-*s  %-*s  ->  %s\n", maxLabel, label, maxOldBase, oldBase, newBase)
		if err != nil {
			return err
		}
	}

	renamed := result.Count(app.StatusRenamed)
	skipped := result.Count(app.StatusSkipped)
	failed := result.Count(app.StatusFailed)

	if result.DryRun {
		renamed = result.Count(app.StatusPlanned)
	}

	_, err := fmt.Fprintf(w, "\n%d renamed, %d skipped, %d failed\n", renamed, skipped, failed)
	return err
}
