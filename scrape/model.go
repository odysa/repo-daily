// Package scrape provides functions for scraping GitHub pulse pages.
package scrape

import "time"

// IssueStatus represents the status of an issue.
type IssueStatus string

// PRStatus represents the status of a pull request.
type PRStatus string

// Issue status constants.
const (
	OpenIssue   IssueStatus = "open"
	ClosedIssue IssueStatus = "closed"
)

// PR status constants.
const (
	OpenPR   PRStatus = "open"
	ClosedPR PRStatus = "closed"
)

// GitHubPulseReport represents a summary report of a GitHub pulse page.
type GitHubPulseReport struct {
	OpenPR          []PR          // Open pull requests.
	ClosedPR        []PR          // Closed pull requests.
	OpenIssues      []Issue       // Open issues.
	ClosedIssues    []Issue       // Closed issues.
	Summary         string        // Summary text.
	TopContributors []Contributor // Top contributors.
}

// PR represents a pull request.
type PR struct {
	Title     string     // Title of the pull request.
	Status    PRStatus   // Status of the pull request.
	Content   string     // Content of the pull request.
	TimeStamp *time.Time // Time stamp of the pull request.
}

// Issue represents an issue.
type Issue struct {
	Title     string      // Title of the issue.
	Status    IssueStatus // Status of the issue.
	Content   string      // Content of the issue.
	TimeStamp *time.Time  // Time stamp of the issue.
}

// Contributor represents a GitHub contributor.
type Contributor struct {
	Name     string `json:"name"`     // Name of the contributor.
	Login    string `json:"login"`    // Login name of the contributor.
	Gravatar string `json:"gravatar"` // Gravatar URL of the contributor.
	Commits  int    `json:"commits"`  // Number of commits by the contributor.
}
