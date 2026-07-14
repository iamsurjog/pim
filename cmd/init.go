/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
		// TODO: Take python version
		// fmt.Printf("Enter python version to use:")
		// fmt.Scan()
		lock := `{
    "version": 3.14,
    "dependencies": [],
	"scripts": [],
}`
		_, err := os.Stat(fileName)
		if err == nil {
			fmt.Printf("File '%s' already exists. Skipping creation.\n", fileName)
			return
		}

		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("File '%s' not found. Creating it now...\n", fileName)

			// 3. Write the file with standard read/write permissions (0644)
			writeErr := os.WriteFile(fileName, []byte(lock), 0644)
			if writeErr != nil {
				fmt.Printf("Failed to write to file: %v\n", writeErr)
				return
			}

			fmt.Println("Successfully wrote config to lock.json!")
		} else {
			// Handle edge cases like permission denied errors when checking the file
			fmt.Printf("An unexpected error occurred while checking the file: %v\n", err)
		}
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
