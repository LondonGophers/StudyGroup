package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type issue struct {
	Title     string   `json:"title"`
	Labels    []string `json:"labels"`
	Status    string   `json:"state"`
	Body      string   `json:"body"`
	CreatedAt string   `json:"created_at"`
}

func readIssue() error {
	fmt.Println("Retrieving the issue..")

	var read = flag.NewFlagSet("read", flag.ExitOnError)

	owner := read.String("owner", "", "repository owner")
	repo := read.String("repo", "", "GitHub repository")
	id := read.String("id", "", "ID of the issue to retrieve")

	if err := read.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Error parsing flags: %v\n", err)
	}

	// Check that the required flags are set
	if *owner == "" || *repo == "" || *id == "" {
		return errors.New("Usage: issuer read -owner OWNER -repo REPO -id ID\n")
	}

	url := fmt.Sprintf("%srepos/%s/%s/issues/%s", baseURL, *owner, *repo, *id)
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error sending http request: %v\n", err)
	}
	if !(isSuccessfulStatus(res.StatusCode)) {
		return fmt.Errorf("Request to Github API failed with status code: %v\n", res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return fmt.Errorf("Error reading http response body: %v\n", err)
	}
	return printIssue(data)
}

func printIssue(b []byte) error {
	i := issue{}
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	fmt.Printf("Issue: %s\nStatus: %s\nDescription: %s\nCreated at: %s\n", i.Title, i.Status, i.Body, i.CreatedAt)
	return nil
}
