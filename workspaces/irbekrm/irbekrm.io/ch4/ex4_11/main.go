package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type createFlags struct {
	Owner     string `json:"-"`
	Repo      string `json:"-"`
	Title     string
	Milestone int    `json:",omitempty"`
	Labels    string `json:",omitempty"`
	Assignees string `json:",omitempty"`
}

// TODO: make it so that a user can pass a non-interactive flag for create (-body flag and no opening of text editor)
var (
	create                                = flag.NewFlagSet("create", flag.ExitOnError)
	owner, repo, title, labels, assignees *string
	milestone                             *int
)

const baseURL = "https://api.github.com/"

func init() {
	owner = create.String("owner", "", "repository owner")
	repo = create.String("repo", "", "GitHub repository")
	title = create.String("title", "", "Title of the issue")
	assignees = create.String("assignee", "", "Comma separated logins for users to assign this issue")
	milestone = create.Int("milestone", 0, "milestone to associate issue with")
	labels = create.String("labels", "", "Comma separated labels to associate with the issue")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: issuer create")
		os.Exit(1)
	}
	switch command := os.Args[1]; command {
	case "create":
		if err := create.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Creating an issue..")
		createIssue()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %v\n", command)
	}
}

// GitHub v3 API https://docs.github.com/en/rest/reference/issues#create-an-issue
func createIssue() {
	cf := &createFlags{Owner: *owner, Title: *title, Repo: *repo, Milestone: *milestone, Labels: *labels, Assignees: *assignees}
	if cf.Owner == "" || cf.Repo == "" || cf.Title == "" {
		fmt.Fprint(os.Stderr, "Usage: issuer create -owner OWNER -title TITLE -repo REPO\n")
		os.Exit(1)
	}
	url := fmt.Sprintf("%srepos/%s/%s/issues", baseURL, cf.Owner, cf.Repo)
	data, err := json.Marshal(cf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling json: %v\n", err)
		os.Exit(1)
	}
	r := bytes.NewReader(data)
	fmt.Println(url)
	resp, err := http.Post(url, "Accept: application/vnd.github.v3+json", r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating an issue: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response: %v\n", string(body))
}
