package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	year  = 31536000
	month = 2626560
)

// go run main go repo:golang/go is:open json decoder
func main() {
	var mo, yo, new []*github.Issue
	now := time.Now().Unix()
	ma, ya := time.Unix(now-month, 0), time.Unix(now-year, 0)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Printf("Total number of issues: %d\n", result.TotalCount)
	printer("Issues that are more than a year old", yo)
	printer("Issues that are more than a month old", mo)
	printer("Issues that are less than a month old", new)
}

func printer(description string, issues []*github.Issue) {
	fmt.Printf("%s: %d\n", description, len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println()
}
