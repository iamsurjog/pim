/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a package dependency from the project configuration.",
	Long: `Strips the specified package dependency from the local project manifest file and
updates the workspace path mapping variables. This completely detaches the dependency
from the dynamic runtime layer without altering or removing packages from the shared
global cache directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
