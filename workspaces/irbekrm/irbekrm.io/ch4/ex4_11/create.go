package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type createData struct {
	Owner     string `json:"-"`
	Repo      string `json:"-"`
	Title     string `json:"title"`
	Milestone int    `json:"milestone,omitempty"`
	Labels    string `json:"labels,omitempty"`
	Assignees string `json:"assignees,omitempty"`
	Body      string `json:"body,omitempty"`
}

// TODO: make it so that a user can pass a non-interactive flag for create (-body flag and no opening of text editor)

// GitHub v3 API https://docs.github.com/en/rest/reference/issues#create-an-issue
func createIssue() error {
	fmt.Println("Creating an issue..")

	var create = flag.NewFlagSet("create", flag.ExitOnError)

	owner := create.String("owner", "", "repository owner")
	repo := create.String("repo", "", "GitHub repository")
	title := create.String("title", "", "Title of the issue")
	assignees := create.String("assignee", "", "Comma separated logins for users to assign this issue")
	milestone := create.Int("milestone", 0, "milestone to associate issue with")
	labels := create.String("labels", "", "Comma separated labels to associate with the issue")
	editor := create.String("editor", "vim", "Text editor that will be opened for user to enter issue description. Choose from vim, vi or nano")

	if err := create.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	cf := &createData{Owner: *owner, Title: *title, Repo: *repo, Milestone: *milestone, Labels: *labels, Assignees: *assignees}

	// Check that the required flags are set
	if cf.Owner == "" || cf.Repo == "" || cf.Title == "" {
		return errors.New("Usage: issuer create -owner OWNER -title TITLE -repo REPO [-assignees ASSIGNEES] [-milestone MILESTONE] [-labels LABELS] [-editor vi|vim|nano]\n")
	}

	// Retrieve Github basic auth creds from environment
	username, password, credsSet := credsFromEnv()
	if !credsSet {
		return errors.New("GITHUB_USERNAME and/or GITHUB_PASSWORD env vars for Github basic auth not set")
	}

	// Check that the editor name is valid
	if !isValidEditor(*editor) {
		return fmt.Errorf("Invalid editor: %v\n", *editor)
	}
	b, err := issueDescription(*editor)
	if err != nil {
		return err
	}
	cf.Body = b

	url := fmt.Sprintf("%srepos/%s/%s/issues", baseURL, cf.Owner, cf.Repo)
	data, err := json.Marshal(cf)
	if err != nil {
		return fmt.Errorf("Error marshaling json: %v\n", err)
	}
	r := bytes.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return fmt.Errorf("Error creating a request: %v\n", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error creating an issue: %v\n", err)
	}
	if !(isSuccessfulStatus(resp.StatusCode)) {
		return fmt.Errorf("Request to Github API failed with status code: %v\n", resp.StatusCode)
	}
	fmt.Printf("Issue %q in repo %q created successfully\n", cf.Title, cf.Repo)
	return nil
}
