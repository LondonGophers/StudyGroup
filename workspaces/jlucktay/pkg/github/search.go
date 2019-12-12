package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	req, errNewReq := http.NewRequest("GET", IssuesSearchURL+"?q="+q, nil)
	if errNewReq != nil {
		return nil, errNewReq
	}
	// Add an HTTP request header indicating that only version 3 of the GitHub API is acceptable.
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, errDo := http.DefaultClient.Do(req)

	if errDo != nil {
		return nil, errDo
	}
	defer resp.Body.Close()

	var result IssuesSearchResult

	if resp.StatusCode == http.StatusUnprocessableEntity {
		return &result, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	if errDecode := json.NewDecoder(resp.Body).Decode(&result); errDecode != nil {
		return nil, errDecode
	}

	return &result, nil
}
