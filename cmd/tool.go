/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// toolCmd represents the tool command
var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Execute an installed package command-line utility directly.",
	Long: `Locates the executable entry-point script generated in the global cache for a
specific package tool (e.g., black, jupyter, flake8) matching your project requirements.
It assembles the necessary environment states on the fly and invokes the utility command
natively, passing along any runtime flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(toolCmd)
}
