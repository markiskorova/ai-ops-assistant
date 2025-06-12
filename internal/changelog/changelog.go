package changelog

import "strings"

type GitCommit struct {
	Message string `json:"message"`
	Author  string `json:"author"`
	Date    string `json:"date"`
}

type ChangelogEntry struct {
	Scope   string `json:"scope"`
	Summary string `json:"summary"`
}

func ParseChangelog(commits []GitCommit) ([]ChangelogEntry, error) {
	var entries []ChangelogEntry

	for _, commit := range commits {
		scope := "general"
		msg := commit.Message

		switch {
		case len(msg) >= 5 && msg[:5] == "feat:":
			scope = "feature"
			msg = msg[5:]
		case len(msg) >= 4 && msg[:4] == "fix:":
			scope = "bugfix"
			msg = msg[4:]
		case len(msg) >= 6 && msg[:6] == "chore:":
			scope = "maintenance"
			msg = msg[6:]
		case len(msg) >= 5 && msg[:5] == "docs:":
			scope = "docs"
			msg = msg[5:]
		}

		entry := ChangelogEntry{
			Scope:   scope,
			Summary: strings.TrimSpace(msg),
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
