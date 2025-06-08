package changelog

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
        entry := ChangelogEntry{
            Scope:   "general", // placeholder, could be inferred from commit message prefix
            Summary: commit.Message,
        }
        entries = append(entries, entry)
    }
    return entries, nil
}
