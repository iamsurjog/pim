/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new package dependency to the active project context.",
	Long: `Resolves a targeted package string via real-time PyPI API metadata, updates the
project's 'project.toml' manifest file, and pulls down the required wheels
concurrently if they are missing from the global cache. It immediately updates the
workspace .env paths to expose the package to your development environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
