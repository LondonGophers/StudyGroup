package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
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

	create.Parse(os.Args[2:])
	if err := create.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	cf := &createData{Owner: *owner, Title: *title, Repo: *repo, Milestone: *milestone, Labels: *labels, Assignees: *assignees}

	// Check that the required flags are set
	if cf.Owner == "" || cf.Repo == "" || cf.Title == "" {
		if err := create.Parse(os.Args[2:]); err != nil {
			return errors.New("Usage: issuer create -owner OWNER -title TITLE -repo REPO [-assignees ASSIGNEES] [-milestone MILESTONE] [-labels LABELS]\n")
		}
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
	// Ask user to enter issue description via the chosen text editor
	tmpfile, err := ioutil.TempFile("", "issue.txt")
	if err != nil {
		return fmt.Errorf("Error creating the description file: %v\n", err)
	}
	defer os.Remove(tmpfile.Name())
	s := []byte("### Please enter issue description below. Do not remove this line. It will not be posted to GitHub. ####\n")
	if _, err := tmpfile.Write(s); err != nil {
		return fmt.Errorf("Error writing to the description file: %v\n", err)
	}
	if err := tmpfile.Close(); err != nil {
		return fmt.Errorf("Error closing the description file: %v\n", err)
	}
	cmd := exec.Command(*editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("Error opening file with editor: %v\n", err)
	}
	contents, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return fmt.Errorf("Error reading file: %v\n", err)
	}
	// Remove the top line with comment
	re := regexp.MustCompile(`####\n((.|\n)+)`)
	c := re.FindStringSubmatch(string(contents))
	if len(c) < 2 {
		return errors.New("Error parsing the description file")
	}
	cf.Body = string(c[1])

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
