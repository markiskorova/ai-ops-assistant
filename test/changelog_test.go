package test

import (
	"ai-ops-assistant/internal/changelog"
	"testing"
)

func TestParseChangelog(t *testing.T) {
	commits := []changelog.GitCommit{
		{Message: "feat: add user login"},
		{Message: "fix: resolve nil pointer"},
		{Message: "docs: update README"},
		{Message: "chore: bump deps"},
		{Message: "refactor DB logic"},
	}

	result, _ := changelog.ParseChangelog(commits)

	expectedScopes := []string{"feature", "bugfix", "docs", "maintenance", "general"}

	for i, entry := range result {
		if entry.Scope != expectedScopes[i] {
			t.Errorf("Expected scope %s, got %s", expectedScopes[i], entry.Scope)
		}
	}
}
