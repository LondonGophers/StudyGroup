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
	Title     string `json:"title"`
	Milestone int    `json:"milestone,omitempty"`
	Labels    string `json:"labels,omitempty"`
	Assignees string `json:"assignees,omitempty"`
	Username  string `json:"-"`
	Password  string `json:"-"`
}

// TODO: make it so that a user can pass a non-interactive flag for create (-body flag and no opening of text editor)
var (
	create                                                    = flag.NewFlagSet("create", flag.ExitOnError)
	owner, repo, title, labels, assignees, username, password *string
	milestone                                                 *int
)

const baseURL = "https://api.github.com/"

func init() {
	owner = create.String("owner", "", "repository owner")
	repo = create.String("repo", "", "GitHub repository")
	title = create.String("title", "", "Title of the issue")
	assignees = create.String("assignee", "", "Comma separated logins for users to assign this issue")
	milestone = create.Int("milestone", 0, "milestone to associate issue with")
	labels = create.String("labels", "", "Comma separated labels to associate with the issue")
	username = create.String("username", "", "Username for basic auth")
	password = create.String("password", "", "password for basic auth")
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
	cf := &createFlags{Owner: *owner, Title: *title, Repo: *repo, Milestone: *milestone, Labels: *labels, Assignees: *assignees, Password: *password, Username: *username}
	if cf.Owner == "" || cf.Repo == "" || cf.Title == "" {
		fmt.Fprint(os.Stderr, "Usage: issuer create -owner OWNER -title TITLE -repo REPO\n")
		os.Exit(1)
	}
	url := fmt.Sprintf("%srepos/%s/%s/issues", baseURL, cf.Owner, cf.Repo)
	fmt.Println(url)
	data, err := json.Marshal(cf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	r := bytes.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating a request: %v\n", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.SetBasicAuth(cf.Username, cf.Password)
	resp, err := client.Do(req)
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
