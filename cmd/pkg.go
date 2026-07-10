/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pkgCmd represents the pkg command
var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Inspect and query PyPI package metadata globally.",
	Long: `The informational command group for querying remote registry metadata, checking
physical package configurations on disk, and executing exploration tools that do not
alter local project source directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(pkgCmd)
}
