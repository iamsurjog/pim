/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a fresh workspace manifest and select a Python version.",
	Long: `Interactively creates a 'pim.json' project configuration file in the current
working directory. It automatically scans your locally cached Python installations,
prompts you to select the active interpreter version via an interactive menu, and
hooks the workspace by writing an optimized local .env file for IDE autocompletion.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Enter python version to use:")
		fmt.Scan()
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
