// Package scrape provides functions for scraping GitHub pulse pages.
package scrape

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// DailyPulseURLTemplate is the template URL for the GitHub pulse page.
const DailyPulseURLTemplate = "%s/pulse/daily"

// Scrape scrapes a GitHub pulse page and returns a *goquery.Document or an error.
// It takes a string parameter, repoURL, which represents the URL of the repository to scrape.
func Scrape(repoURL string) (*goquery.Document, error) {
	url := fmt.Sprintf(DailyPulseURLTemplate, repoURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("failed to close resp body %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request error code %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Parse parses a *goquery.Document and returns a *GitHubPulseReport or an error.
func Parse(doc *goquery.Document) (*GitHubPulseReport, error) {
	contributors, err := parseContributors(doc)
	if err != nil {
		return nil, err
	}

	openIssues, err := parseOpenIssues(doc)
	if err != nil {
		return nil, err
	}

	closedIssues, err := parseClosedIssues(doc)
	if err != nil {
		return nil, err
	}

	openPRs, err := parseOpenPRs(doc)
	if err != nil {
		return nil, err
	}

	closedPRs, err := parseClosedPRs(doc)
	if err != nil {
		return nil, err
	}

	return &GitHubPulseReport{
		TopContributors: contributors,
		OpenIssues:      openIssues,
		ClosedIssues:    closedIssues,
		OpenPR:          openPRs,
		ClosedPR:        closedPRs,
	}, nil
}

// parseContributors parses the top contributors from a *goquery.Document and returns a slice of Contributors or an error.
func parseContributors(doc *goquery.Document) ([]Contributor, error) {
	// to do
	return nil, nil
}

// parseOpenIssues parses the open issues from a *goquery.Document and returns a slice of Issues or an error.
func parseOpenIssues(doc *goquery.Document) ([]Issue, error) {
	// to do
	return parseIssues(doc, true)
}

// parseClosedIssues parses the closed issues from a *goquery.Document and returns a slice of Issues or an error.
func parseClosedIssues(doc *goquery.Document) ([]Issue, error) {
	return parseIssues(doc, false)
}

// parseIssues parses the issues from a *goquery.Document and returns a slice of Issues or an error.
func parseIssues(doc *goquery.Document, open bool) ([]Issue, error) {
	// to do
	return nil, nil
}

// parseOpenPRs parses the open pull requests from a *goquery.Document and returns a slice of PRs or an error.
func parseOpenPRs(doc *goquery.Document) ([]PR, error) {
	return parsePRs(doc, true)
}

// parseClosedPRs parses the closed pull requests from a *goquery.Document and returns a slice of PRs or an error.
func parseClosedPRs(doc *goquery.Document) ([]PR, error) {
	return parsePRs(doc, false)
}

// parsePRs parses the pull requests from a *goquery.Document and returns a slice of PRs or an error.
func parsePRs(doc *goquery.Document, open bool) ([]PR, error) {
	// to do
	return nil, nil
}
