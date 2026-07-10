/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command (under pkg)
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Perform a rapid local fuzzy search against all PyPI package names.",
	Long: `Runs a high-performance Levenshtein distance matching sequence entirely within system
memory against a locally cached copy of the PEP 691 PyPI simple index. It prints a
status header warning showing data freshness and exposes a manual '--refresh' override
flag to update the download layout.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	pkgCmd.AddCommand(searchCmd)
}
