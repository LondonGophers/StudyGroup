package github

import "time"

// DEVELOPMENT OF THIS HAS NOW MOVED TO https://github.com/betandr/prt
// THIS WON'T BE UPDATED ANYMORE SO WILL CONTAIN BUGS ETC

// PullRequestsURL is the Github API URL
const PullRequestsURL = "https://api.github.com"

// PullRequestsResult represents a collection of PullRequests
type PullRequestsResult struct {
	PullRequests []*PullRequest
}

// User represents a GitHub user
type User struct {
	Login string
}

// Head is the branch the PR is on
type Head struct {
	Ref string
}

// Base is the branch this PR's branch is from
type Base struct {
	Ref string
}

// Users is collection of `User`s
type Users []User

// CommitDetails are extra details for the commit
type CommitDetails struct {
	Message string
}

// Commit from a PR
type Commit struct {
	SHA           string
	Author        User
	CommitDetails `json:"commit"`
	UpdatedAt     string `json:"updated_at"`
}

// Commits is a collection of `Commit`s (obvs)
type Commits []Commit

// Comment from a PR
type Comment struct {
	ID        int
	User      User
	UpdatedAt time.Time `json:"updated_at"`
	Body      string
}

// Status from a PR
type Status struct {
	ID          int
	State       string
	Description string
	TargetURL   string    `json:"target_url"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewPullRequest is used when opening a new pull request
type NewPullRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Head  string `json:"head"`
	Base  string `json:"base"`
}

// PatchPullRequest is used to update a PR
type PatchPullRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
	Base  string `json:"base"`
}

// Error shows feedback from the API
type Error struct {
	Message string
	Errors  []*ErrorMessage
}

// ErrorMessage is used in an `Error`
type ErrorMessage struct {
	Message string
}

// MergeStatus shows the status of a merge request
type MergeStatus struct {
	SHA     string `json:"sha"`
	Merged  bool   `json:"merged"`
	Message string `json:"message"`
}

// PullRequest represents a single pull request
type PullRequest struct {
	Number             int
	State              string
	Locked             bool
	Title              string
	Body               string
	StatusesURL        string `json:"statuses_url"`
	HTMLUrl            string `json:"html_url"`
	User               User
	CreatedAt          time.Time `json:"created_at"`
	MergedAt           time.Time `json:"merged_at"`
	MergeCommitSha     string    `json:"merge_commit_sha"`
	Head               Head
	Base               Base
	RequestedReviewers Users `json:"requested_reviewers"`
	Mergeable          bool
	Rebaseable         bool
	MergeableState     string `json:"mergeable_state"`
	MergedBy           User   `json:"merged_by"`
	Commits            int
	Additions          int
	Deletions          int
	ChangedFiles       int `json:"changed_files"`
}
