package main

import (
	"bytes"
	"testing"
	"time"

	"gopl.io/ch4/github"
)

func TestCategorizeIssues(t *testing.T) {
	tests := []struct {
		name        string
		issueGetter func([]string) *github.IssuesSearchResult
		wantOut     string
	}{
		{
			name: "Success- finds issues created at different times",
			issueGetter: func(s []string) *github.IssuesSearchResult {
				yo := &github.Issue{CreatedAt: time.Unix(time.Now().Unix()-Year*2, 0), Title: "More than a year old", User: &github.User{}}
				mo := &github.Issue{CreatedAt: time.Unix(time.Now().Unix()-Month*2, 0), Title: "More than a month old", User: &github.User{}}
				new := &github.Issue{CreatedAt: time.Unix(time.Now().Unix(), 0), Title: "New", User: &github.User{}}
				return &github.IssuesSearchResult{TotalCount: 3, Items: []*github.Issue{yo, mo, new}}
			},
			wantOut: "Total number of issues: 3\nIssues that are more than a year old: 1\n#0               More than a year old\n\nIssues that are more than a month old: 1\n#0               More than a month old\n\nIssues that are less than a month old: 1\n#0               New\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			CategorizeIssues([]string{}, out, tt.issueGetter)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("CategorizeIssues() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
