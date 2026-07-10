/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pyCmd represents the py command
var pyCmd = &cobra.Command{
	Use:   "py",
	Short: "Manage standalone Python interpreter installations.",
	Long: `The base command tree for inspecting, acquiring, and switching between pre-compiled
CPython interpreter runtimes. It maps directory targets, coordinates downloads, and
isolates standard system python setups from your developer workspaces.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(pyCmd)
}
