package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	Year  = 31536000
	Month = 2626560
)

// go run main.go repo:golang/go is:open json decoder
func main() {
	CategorizeIssues(os.Args[1:], os.Stdout, issueGetter)
}

// CategorizeIssues accepts GitHub search terms, and a function to call GitHub API.
// It categorizes returned GitHub issues into more than year old, more than month old and less than month old.
// Writes the results to io.Writer argument
func CategorizeIssues(terms []string, out io.Writer, issueGetter func([]string) *github.IssuesSearchResult) {
	var mo, yo, new []*github.Issue
	now := time.Now().Unix()
	ma, ya := time.Unix(now-Month, 0), time.Unix(now-Year, 0)
	result := issueGetter(terms)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.Before(ya):
			yo = append(yo, item)
		case item.CreatedAt.Before(ma):
			mo = append(mo, item)
		default:
			new = append(new, item)
		}
	}
	fmt.Fprintf(out, "Total number of issues: %d\n", result.TotalCount)
	printer("Issues that are more than a year old", yo, out)
	printer("Issues that are more than a month old", mo, out)
	printer("Issues that are less than a month old", new, out)
}

func issueGetter(terms []string) *github.IssuesSearchResult {
	res, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func printer(description string, issues []*github.Issue, out io.Writer) {
	fmt.Fprintf(out, "%s: %d\n", description, len(issues))
	for _, item := range issues {
		fmt.Fprintf(out, "#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Fprintln(out)
}
