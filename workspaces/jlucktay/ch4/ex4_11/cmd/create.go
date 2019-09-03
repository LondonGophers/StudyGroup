package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Flags
var fTitle, fComment *string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create GitHub issues from the command line",
	Long:  `Create a new GitHub issue from the command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\"%s --title '%s' --comment '%s'\" called.\n", cmd.CommandPath(), *fTitle, *fComment)
		create(*fTitle, *fComment)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	fTitle = createCmd.Flags().StringP("title", "t", "", "The title line")
	fComment = createCmd.Flags().StringP("comment", "c", "", "The comment body")
}

func create(title, comment string) {
	fmt.Printf("Creating an issue with title '%s' and comment '%s'!\n", title, comment)
}
