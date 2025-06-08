package changelog

import (
    "testing"
)

func TestParseChangelog(t *testing.T) {
    commits := []GitCommit{
        {Message: "fix: corrected login bug", Author: "dev", Date: "2025-06-01"},
    }

    changelog, err := ParseChangelog(commits)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if len(changelog) == 0 {
        t.Errorf("expected at least one changelog entry")
    }
}
