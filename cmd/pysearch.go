/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pySearchCmd represents the py search command
var pySearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for available remote CPython runtime versions.",
	Long: `Queries a cached index of pre-compiled, portable standalone CPython binaries hosted
upstream (via python-build-standalone). It prints a flat, newline-separated text
stream optimized for piping directly into UNIX tools like fzf or ripgrep.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	pyCmd.AddCommand(pySearchCmd)
}
