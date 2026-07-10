/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List all Python interpreter versions downloaded locally.",
	Long: `Scans the localized storage partition directory where pim maintains its standalone
runtimes. It outputs a clean, unformatted plain-text list of version numbers, making
it instantly compatible with shell scripts, xargs loops, and interactive fuzzy finders.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	pyCmd.AddCommand(versionsCmd)
}
