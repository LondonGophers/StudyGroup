package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/LondonGophers/StudyGroup/workspaces/jlucktay/pkg/github"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Search GitHub issues from the command line",
	Long: `Search for GitHub issues from the command line.
All arguments are passed through to GitHub search.

For example:

$ ghissues read author:foo language:bar comments:>50

Further reading on search syntax:
- https://help.github.com/articles/searching-issues-and-pull-requests
- https://help.github.com/articles/understanding-the-search-syntax
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("'%s %s' called.\n", cmd.CommandPath(), strings.Join(args, " "))
		read(args)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	//
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")
	//
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(readCmd)
}

func read(searchTerms []string) {
	issues, err := github.SearchIssues(searchTerms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n\n", issues.TotalCount)

	for _, issue := range issues.Items {
		fmt.Printf("[%6d] %-12s: %s\n         %s\n\n", issue.Number, issue.User.Login, issue.Title, issue.HTMLURL)
	}
}
