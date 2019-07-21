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

	reqURL := IssuesURL + "?q=" + q

	req, err := http.NewRequest("GET", reqURL, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	acceptHeaders := []string{"application/vnd.github.symmetra-preview+json", "application/vnd.github.sailor-v-preview+json"}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("Error: %s", res.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return &result, nil
}
