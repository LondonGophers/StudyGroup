package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateIssue creates a new issue under the designated owner/repo.
func CreateIssue(create IssueCreate, auth Auth, owner, repo string) (*IssueCreateResult, error) {
	reqBodyJSON, err := json.Marshal(create)
	if err != nil {
		return nil, err
	}

	createURL := fmt.Sprintf(IssueCreateURL, owner, repo)

	req, err := http.NewRequest("POST", createURL, strings.NewReader(string(reqBodyJSON)))
	req.SetBasicAuth(auth.Username, auth.Password)

	if err != nil {
		return nil, err
	}
	// Add an HTTP request header indicating that only version 3 of the GitHub API is acceptable.
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create request failed: %s", resp.Status)
	}

	var result IssueCreateResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
