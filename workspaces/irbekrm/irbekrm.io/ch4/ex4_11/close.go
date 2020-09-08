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

type closeData struct {
	Id    string `json:"id"`
	Owner string `json:"-"`
	Repo  string `json:"-"`
	State string `json:"state"`
}

func closeIssue() error {
	fmt.Println("Closing the issue..")

	var close = flag.NewFlagSet("close", flag.ExitOnError)

	id := close.String("id", "", "id of the issue to close")
	owner := close.String("owner", "", "repository owner")
	repo := close.String("repo", "", "GitHub repository")

	if err := close.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	clf := &closeData{Id: *id, Owner: *owner, Repo: *repo}
	clf.State = "closed"

	// Check that the required flags are set
	if clf.Owner == "" || clf.Repo == "" || clf.Id == "" {
		return errors.New("Usage: issuer close -id ID -owner OWNER -repo REPO\n")
	}

	// Retrieve Github basic auth creds from environment
	username, password, credsSet := credsFromEnv()
	if !credsSet {
		return errors.New("GITHUB_USERNAME and/or GITHUB_PASSWORD env vars for Github basic auth not set")
	}

	url := fmt.Sprintf("%srepos/%s/%s/issues/%s", baseURL, clf.Owner, clf.Repo, clf.Id)
	data, err := json.Marshal(clf)
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
		return fmt.Errorf("Error closing the issue: %v\n", err)
	}
	if !(isSuccessfulStatus(resp.StatusCode)) {
		return fmt.Errorf("Request to Github API failed with status code: %v\n", resp.StatusCode)
	}
	fmt.Printf("Issue %q in repo %q closed successfully\n", clf.Id, clf.Repo)
	return nil
}
