// Create a web server that queries GitHub once and then allows navigation of
// the list of bug reports, milestones, and users.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"andr.io/ch4/ex4_14/github"
)

func main() {
	issues := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		terms := r.Form["q"]
		if len(terms) < 1 {
			fmt.Fprint(w, "Please supply a search query, such as ?q=foo")
		} else {
			issues(terms, w)
		}
	}

	http.HandleFunc("/issues/", issues)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
	return
}

func issues(terms []string, out io.Writer) {
	templ := `
	<h1>{{.TotalCount}} issues</h1>
    <table>
    <tr style='text-align: left'>
      <th>#</th>
      <th>State</th>
      <th>User</th>
      <th>Title</th>
	  <th>Milestone</th>
    </tr>
    {{range .Items}}
    <tr>
      <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
      <td>{{.State}}</td>
      <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
      <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
    </tr>
    {{end}}
    </table>
    `

	var issueList = template.Must(template.New("issuelist").Parse(templ))

	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}
