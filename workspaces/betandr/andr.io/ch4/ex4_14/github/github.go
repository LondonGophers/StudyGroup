package github

import "time"

// IssuesURL is the GitHub API URL for bug reports / issues
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Milestone struct {
	Number  int
	Title   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Milestone Milestone
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
