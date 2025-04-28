/*
Copyright © 2025 Ahmed Mohamed <ahmedmohamed24.dev@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mentorship-crud-app",
	Short: "Mentorship CRUD API",
	Long:  `This project handles a Documenting System, where each document has the following fields: ID, Title, Author, Content, Created_at, and Updated_at.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
