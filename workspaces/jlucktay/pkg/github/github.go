// Package github provides a Go API for the GitHub issue tracker.
// See: https://developer.github.com/v3/search/#search-issues
package github

import "time"

const APIPrefix = "https://api.github.com"
const IssuesSearchURL = APIPrefix + "/search/issues"
const IssueCreateURL = APIPrefix + "/repos/%s/%s/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*IssueSearchResult
}

type IssueSearchResult struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time  `json:"created_at"`
	Body      string     // in Markdown format
	Milestone *Milestone `json:"milestone"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueCreate struct {
	Title string `json:"title"` // The title of the issue.
	Body  string `json:"body"`  // The contents of the issue.
}

type IssueCreateResult struct {
	HTMLURL string `json:"html_url"`
}

type Auth struct {
	Username, Password string
}

type Milestone struct {
	HTMLURL string `json:"html_url"`
	Title   string `json:"title"`
	Creator *User  `json:"creator"`
}
