package test

import (
	"testing"

	"ai-ops-assistant/internal/changelog"
	"github.com/stretchr/testify/assert"
)

func TestParseChangelog(t *testing.T) {
	commits := []changelog.GitCommit{
		{Message: "feat: add login feature"},
		{Message: "fix: resolve crash on save"},
		{Message: "refactor: cleanup code"},
	}

	entries, err := changelog.ParseChangelog(commits)
	assert.NoError(t, err)
	assert.Len(t, entries, 3)
}
