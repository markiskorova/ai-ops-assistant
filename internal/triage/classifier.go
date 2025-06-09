package triage

import "strings"

func containsCriticalKeywords(text string) bool {
	return containsAny(text, []string{"panic", "outage", "down", "failure"})
}

func containsDatabaseKeywords(text string) bool {
	return containsAny(text, []string{"db", "database", "query", "sql"})
}

func containsAny(text string, keywords []string) bool {
	for _, word := range keywords {
		if strings.Contains(strings.ToLower(text), strings.ToLower(word)) {
			return true
		}
	}
	return false
}
