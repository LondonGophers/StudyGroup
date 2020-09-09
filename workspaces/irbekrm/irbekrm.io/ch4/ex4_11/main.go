package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

const baseURL = "https://api.github.com/"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: issuer COMMAND FLAGS")
		os.Exit(1)
	}
	switch command := os.Args[1]; command {
	case "create":
		if err := createIssue(); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating an issue: %v\n", err)
			os.Exit(1)
		}
	case "read":
		if err := readIssue(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the issue: %v\n", err)
			os.Exit(1)
		}
	case "update":
		if err := updateIssue(); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating the issue: %v\n", err)
			os.Exit(1)
		}
	case "close":
		if err := closeIssue(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing the issue: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %v\n", command)
	}
}

func isSuccessfulStatus(s int) bool {
	return s >= 200 && s <= 299
}

func credsFromEnv() (string, string, bool) {
	username, uSet := os.LookupEnv("GITHUB_USERNAME")
	password, pSet := os.LookupEnv("GITHUB_PASSWORD")
	return username, password, uSet && pSet
}

func isValidEditor(e string) bool {
	if e == "vim" || e == "vi" || e == "nano" {
		return true
	}
	return false
}

// Ask user to enter issue description via the chosen text editor
func issueDescription(e string) (string, error) {
	tmpfile, err := ioutil.TempFile("", "issue.txt")
	if err != nil {
		return "", fmt.Errorf("Error creating the description file: %v\n", err)
	}
	defer os.Remove(tmpfile.Name())
	s := []byte("### Please enter issue description below. Do not remove this line. It will not be posted to GitHub. ####\n")
	if _, err := tmpfile.Write(s); err != nil {
		return "", fmt.Errorf("Error writing to the description file: %v\n", err)
	}
	if err := tmpfile.Close(); err != nil {
		return "", fmt.Errorf("Error closing the description file: %v\n", err)
	}
	cmd := exec.Command(e, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil {
		return "", fmt.Errorf("Error opening file with editor: %v\n", err)
	}
	contents, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v\n", err)
	}
	// Remove the top line with comment
	re := regexp.MustCompile(`####\n((.|\n)+)`)
	c := re.FindStringSubmatch(string(contents))
	if len(c) < 2 {
		return "", errors.New("Error parsing the description file")
	}
	return c[1], nil
}
