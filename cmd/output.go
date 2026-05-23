package cmd

import (
	"fmt"
	"io"

	"github.com/shricodev/case/internal/app"
)

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

func printResult(w io.Writer, result app.Result) error {
	for _, item := range result.Items {
		msg := messageForStatus(item.Status)

		if item.Status == app.StatusFailed && item.Error != nil {
			_, err := fmt.Fprintf(w, "%s  %s: %v\n", msg, item.OldPath, item.Error)
			if err != nil {
				return err
			}

			continue
		}

		_, err := fmt.Fprintf(w, "%s %s -> %s\n", msg, item.OldPath, item.NewPath)
		if err != nil {
			return err
		}
	}

	return nil
}
