// Modify `issues` to report the results in age categories, say less than a month old, less than a year old, and more
// than a year old.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	var monthOld, yearOld, older []*github.Issue

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		switch {
		case time.Since(item.CreatedAt) < time.Hour*24*30:
			monthOld = append(monthOld, item)
		case time.Since(item.CreatedAt) < time.Hour*24*365:
			yearOld = append(yearOld, item)
		default:
			older = append(older, item)
		}
	}

	fmt.Printf("Less than a month old:\n%s\n", toString(monthOld))
	fmt.Printf("Less than a year old:\n%s\n", toString(yearOld))
	fmt.Printf("More than a year old:\n%s\n", toString(older))
}

func toString(ghi []*github.Issue) (s string) {
	for _, item := range ghi {
		s += fmt.Sprintf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	return
}
