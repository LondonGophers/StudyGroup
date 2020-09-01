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

type createFlags struct {
	Owner     string `json:"-"`
	Repo      string `json:"-"`
	Title     string `json:"title"`
	Milestone int    `json:"milestone,omitempty"`
	Labels    string `json:"labels,omitempty"`
	Assignees string `json:"assignees,omitempty"`
	Username  string `json:"-"`
	Password  string `json:"-"`
}

// TODO: make it so that a user can pass a non-interactive flag for create (-body flag and no opening of text editor)

// GitHub v3 API https://docs.github.com/en/rest/reference/issues#create-an-issue
func createIssue() error {
	fmt.Println("Creating an issue..")
	var create = flag.NewFlagSet("create", flag.ExitOnError)
	var owner, repo, title, labels, assignees, username, password *string
	var milestone *int
	owner = create.String("owner", "", "repository owner")
	repo = create.String("repo", "", "GitHub repository")
	title = create.String("title", "", "Title of the issue")
	assignees = create.String("assignee", "", "Comma separated logins for users to assign this issue")
	milestone = create.Int("milestone", 0, "milestone to associate issue with")
	labels = create.String("labels", "", "Comma separated labels to associate with the issue")
	username = create.String("username", "", "Username for basic auth")
	password = create.String("password", "", "password for basic auth")
	create.Parse(os.Args[2:])
	if err := create.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}
	cf := &createFlags{Owner: *owner, Title: *title, Repo: *repo, Milestone: *milestone, Labels: *labels, Assignees: *assignees, Password: *password, Username: *username}
	if cf.Owner == "" || cf.Repo == "" || cf.Title == "" {
		if err := create.Parse(os.Args[2:]); err != nil {
			return errors.New("Usage: issuer create -owner OWNER -title TITLE -repo REPO\n")
		}
	}
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
	req.SetBasicAuth(cf.Username, cf.Password)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error creating an issue: %v\n", err)
	}
	if !(statusIsSuccessful(resp.StatusCode)) {
		return fmt.Errorf("Non-successful status code: %v\n", resp.StatusCode)
	}
	fmt.Printf("Issue %q in repo %q created successfully\n", cf.Title, cf.Repo)
	return nil
}
