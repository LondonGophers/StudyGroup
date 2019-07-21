// Build a tool that lets users create, read, update, and close GitHub i̶s̶s̶u̶e̶s̶
// pull requests from the command line, invoking their preferred text editor
// when substantial text input is required.
package main

// DEVELOPMENT OF THIS HAS NOW MOVED TO https://github.com/betandr/prt
// THIS WON'T BE UPDATED ANYMORE SO WILL CONTAIN BUGS ETC

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"andr.io/ch4/ex4_11/github"
)

func locked(b bool) string {
	if b {
		return "locked"
	}

	return "unlocked"
}

func daysAgo(then time.Time) int {
	now := time.Now()
	days := now.Sub(then).Hours() / 24
	return int(days)
}

func merged(at time.Time) bool {
	if at.Year() != 1 {
		return true
	}

	return false
}

func makePhrase(inverse bool, s string) string {
	if inverse {
		return s
	}

	return "not " + s
}

func renderCommit(commit *github.Commit) {
	fmt.Printf("%s\n", strings.Replace(commit.Message, "\n", " ", -1))
	fmt.Printf("%s [\x1b[36;1m%s\x1b[0m]\n", commit.Author.Login, commit.SHA)
}

func renderComment(comment *github.Comment) {
	fmt.Printf("%s\n",
		strings.Replace(comment.Body, "\n", " ", -1))

	fmt.Printf("%s [\x1b[31;1m%d\x1b[0m] %d day(s) ago\n",
		comment.User.Login,
		comment.ID,
		daysAgo(comment.UpdatedAt))
}

func renderStatus(status *github.Status) {
	fmt.Printf("\x1b[33;1m%s\x1b[0m: %s %d day(s) ago\n",
		status.State,
		status.Description,
		daysAgo(status.UpdatedAt))
}

// renderPull renders the pull request `pr` in a short form if `extra` is false
// or with extra fields if `extra` is true.
func renderPull(pr *github.PullRequest, extra bool) {
	fmt.Printf("[\x1b[31;1m%d\x1b[0m] %s (\x1b[33;1m%s\x1b[0m", pr.Number, pr.Title, pr.State)
	if extra {
		fmt.Printf("|\x1b[33;1m%s\x1b[0m|\x1b[33;1m%s\x1b[0m",
			makePhrase(pr.Mergeable, "mergable"),
			makePhrase(pr.Rebaseable, "rebaseable"))

		if pr.MergeableState != "unknown" {
			fmt.Printf("|\x1b[33;1m%s\x1b[0m", pr.MergeableState)
		}
	}
	fmt.Println(")")
	fmt.Printf("Branch: \x1b[32;1m%s\x1b[0m Base: \x1b[32;1m%s\x1b[0m ", pr.Head.Ref, pr.Base.Ref)
	fmt.Printf("SHA: \x1b[36;1m%s\x1b[0m\n", pr.MergeCommitSha)

	if extra {
		fmt.Printf("%d commit(s), %d addition(s), %d deletion(s), %d changed file(s)\n",
			pr.Commits,
			pr.Additions,
			pr.Deletions,
			pr.ChangedFiles)
	}

	if merged(pr.MergedAt) {
		fmt.Printf("Raised by %s and merged by %s %d days ago\n", pr.User.Login, pr.MergedBy.Login, daysAgo(pr.MergedAt))
	} else {
		fmt.Printf("Raised by %s %d days ago\n", pr.User.Login, daysAgo(pr.CreatedAt))
	}

	if len(pr.RequestedReviewers) > 0 {
		rs := []string{}
		for _, r := range pr.RequestedReviewers {
			rs = append(rs, r.Login)
		}
		fmt.Printf("Requested reviewers: %s\n", strings.Join(rs, ", "))
	}

	if extra {
		if len(pr.Body) > 0 {
			fmt.Printf("%s\n\n", pr.Body)
		} else {
			fmt.Println("")
		}
	}
}

func renderMergeStatus(status *github.MergeStatus) {
	fmt.Printf("%s [\x1b[36;1m%s\x1b[0m]\n", status.Message, status.SHA)
}

func main() {
	// This is a work in progress; it's not finished yet and pretty brittle! ;)

	if len(os.Args) <= 1 {
		fmt.Println("Usage: prt {list|get|create|delete} {owner/repo} {...}")
		os.Exit(0)
	}

	if os.Args[1] == "help" {
		fmt.Println("List:\tprt list owner/repo (optional: --all)")
		fmt.Println("Get:\tprt get owner/repo {number}")
		fmt.Println("Create:\tprt create owner/repo {branch} {base} {title}")
		fmt.Println("Update:\tprt update owner/repo {number} {base} {title}")
		fmt.Println("Merge:\tprt merge owner/repo {number} (optional: {title} {message} {method} (merge, squash or rebase))")
		fmt.Println("Close:\tprt close owner/repo {number}")
		os.Exit(0)
	}

	if os.Args[1] == "update" {
		number, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not update PR with number %s", os.Args[3])
			os.Exit(1)
		}

		pr, err := github.GetPullRequest(os.Args[2], number)
		if err != nil {
			log.Fatal(err)
		}

		filename := "4afc7c1fecb812c8cb140d072315a8a5"
		writeErr := ioutil.WriteFile(filename, []byte(pr.Body), 0644)
		if writeErr != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("vim", filename)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()

		body, _ := ioutil.ReadFile(filename)

		os.Remove(filename)

		updatedPr, err := github.UpdatePullRequest(
			os.Args[2],
			number,
			os.Args[5],
			string(body),
			pr.State,
			os.Args[4])
		if err != nil {
			log.Fatal(err)
		}

		renderPull(updatedPr, true)

	} else if os.Args[1] == "close" {
		number, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not update PR with number %s", os.Args[3])
			os.Exit(1)
		}

		pr, err := github.GetPullRequest(os.Args[2], number)
		if err != nil {
			log.Fatal(err)
		}

		updatedPr, err := github.UpdatePullRequest(
			os.Args[2],
			number,
			pr.Title,
			pr.Body,
			"closed",
			pr.Base.Ref)
		if err != nil {
			log.Fatal(err)
		}

		renderPull(updatedPr, true)

	} else if os.Args[1] == "merge" {
		var title string
		var message string
		var method string

		if len(os.Args) > 4 && len(os.Args[4]) > 0 {
			title = os.Args[4]
		}

		if len(os.Args) > 5 && len(os.Args[5]) > 0 {
			message = os.Args[5]
		}

		if len(os.Args) > 6 && len(os.Args[6]) > 0 {
			method = os.Args[6]
		} else {
			method = "merge"
		}

		number, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not merge PR with number %s", os.Args[3])
			os.Exit(1)
		}

		status, err := github.MergePullRequest(os.Args[2], number, title, message, method)
		if err != nil {
			log.Fatal(err)
		}

		renderMergeStatus(status)

	} else if os.Args[1] == "create" {
		filename := "4afc7c1fecb812c8cb140d072315a8a5"
		cmd := exec.Command("vim", filename)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()

		body, _ := ioutil.ReadFile(filename)

		os.Remove(filename)

		pr, err := github.CreatePullRequest(os.Args[2], os.Args[5], string(body), os.Args[3], os.Args[4])
		if err != nil {
			log.Fatal(err)
		}

		renderPull(pr, false)

	} else if os.Args[1] == "get" {
		number, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get PR with number %s from %s", os.Args[4], os.Args[3])
			os.Exit(1)
		}

		pr, err := github.GetPullRequest(os.Args[2], number)
		if err != nil {
			log.Fatal(err)
		}

		commits, err := github.ListPullRequestCommits(os.Args[2], number)
		if err != nil {
			log.Fatal(err)
		}

		comments, err := github.ListPullRequestComments(os.Args[2], number)
		if err != nil {
			log.Fatal(err)
		}

		statuses, err := github.ListPullRequestStatuses(pr.StatusesURL)
		if err != nil {
			log.Fatal(err)
		}

		renderPull(pr, true)

		if len(statuses) > 0 {
			fmt.Println("Status:")
			renderStatus(statuses[0])
		}

		if len(commits) > 0 {
			fmt.Println("\nCommits:")
			for _, c := range commits {
				renderCommit(c)
			}
		}

		if len(comments) > 0 {
			fmt.Println("\nComments:")
			for _, c := range comments {
				renderComment(c)
			}
		}
	} else if os.Args[1] == "list" {
		allIssues := false
		if len(os.Args) > 3 && os.Args[3] == "--all" {
			allIssues = true
		}
		result, err := github.ListPullRequests(os.Args[2], allIssues)
		if err != nil {
			log.Fatal(err)
		}

		for _, pr := range result.PullRequests {
			renderPull(pr, false)
		}
	} else {
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
