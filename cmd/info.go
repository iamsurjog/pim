/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display comprehensive metadata and caching status for a package.",
	Long: `Hits the real-time PyPI JSON API to pull down deep package data (including authors,
licenses, and homepages). It cross-references this with your storage cache to tell
you exactly which versions are downloaded locally and which drive volumes they reside on.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	pkgCmd.AddCommand(infoCmd)
}
