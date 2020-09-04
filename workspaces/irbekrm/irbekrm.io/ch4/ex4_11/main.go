package main

import (
	"fmt"
	"os"
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
