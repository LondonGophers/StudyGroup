package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/LondonGophers/StudyGroup/workspaces/jlucktay/pkg/github"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HandleFunc: %s\n", r.URL)
		searchGitHub(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchGitHub(out io.Writer, r *http.Request) {
	if errParse := r.ParseForm(); errParse != nil {
		panic(errParse)
	}

	var ghisr []*github.IssueSearchResult

	searchTerms := r.Form["q"]

	if len(searchTerms) > 0 {
		ghisr = search(searchTerms)
	}

	if errExec := template.Must(template.ParseFiles("get.gohtml")).Execute(out, ghisr); errExec != nil {
		panic(errExec)
	}
}

func search(searchTerms []string) []*github.IssueSearchResult {
	issues, err := github.SearchIssues(searchTerms)
	if err != nil {
		panic(err)
	}

	return issues.Items
}
