package main

type IssueStatus string
type PRStatus string

const (
	OPEN_ISSUE   IssueStatus = "open"
	CLOSED_ISSUE IssueStatus = "closed"
	OPEN_PR      PRStatus    = "open"
	CLOSED_PR    PRStatus    = "closed"
)

type GitHubPulse struct {
	OpenPR          []PR
	ClosedPR        []PR
	OpenIssues      []Issue
	ClosedIssues    []Issue
	Summary         string
	TopContributors []Contributor
}

type PR struct {
	Title   string
	Status  PRStatus
	Content string
}

type Issue struct {
}

type Contributor struct{}
