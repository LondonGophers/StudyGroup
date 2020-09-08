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

type updateData struct {
	Id        string `json:"id"`
	Owner     string `json:"-"`
	Repo      string `json:"-"`
	Title     string `json:"title"`
	Milestone int    `json:"milestone,omitempty"`
	Labels    string `json:"labels,omitempty"`
	Assignees string `json:"assignees,omitempty"`
	Body      string `json:"body,omitempty"`
}

func updateIssue() error {
	fmt.Println("Updating the issue..")

	var update = flag.NewFlagSet("update", flag.ExitOnError)

	id := update.String("id", "", "id of the issue to update")
	owner := update.String("owner", "", "repository owner")
	repo := update.String("repo", "", "GitHub repository")
	title := update.String("title", "", "Title of the issue")
	assignees := update.String("assignee", "", "Comma separated logins for users to assign this issue")
	milestone := update.Int("milestone", 0, "milestone to associate issue with")
	labels := update.String("labels", "", "Comma separated labels to associate with the issue")
	editor := update.String("editor", "vim", "Text editor that will be opened for user to enter issue description. Choose from vim, vi or nano")

	if err := update.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	uf := &updateData{Id: *id, Owner: *owner, Title: *title, Repo: *repo, Milestone: *milestone, Labels: *labels, Assignees: *assignees}

	// Check that the required flags are set
	if uf.Owner == "" || uf.Repo == "" || uf.Id == "" {
		return errors.New("Usage: issuer update -id ID -owner OWNER -repo REPO [-title TITLE] [-assignees ASSIGNEES] [-milestone MILESTONE] [-labels LABELS] [-editor vi|vim|nano]\n")
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
	uf.Body = string(c[1])

	url := fmt.Sprintf("%srepos/%s/%s/issues/%s", baseURL, uf.Owner, uf.Repo, uf.Id)
	data, err := json.Marshal(uf)
	if err != nil {
		return fmt.Errorf("Error marshaling json: %v\n", err)
	}
	r := bytes.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, r)
	if err != nil {
		return fmt.Errorf("Error creating a request: %v\n", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error updating the issue: %v\n", err)
	}
	if !(isSuccessfulStatus(resp.StatusCode)) {
		return fmt.Errorf("Request to Github API failed with status code: %v\n", resp.StatusCode)
	}
	fmt.Printf("Issue %q in repo %q updated successfully\n", uf.Id, uf.Repo)
	return nil
}
