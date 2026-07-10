/*
Copyright © 2026 Sujatro Ganguli iamsurjog@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pim",
	Short: "A high-performance, environment-less Python runtime and dependency orchestrator written in Go",
	Long: `pim (Pip Improved) is a modern toolchain designed to overhaul Python environment 
and package management. By completely abandoning the traditional, slow, and disk-heavy 
virtual environment folder layout, pim manages runtimes dynamically. 

Packages are stored exactly once in a global, drive-localized cache and injected directly 
into the Python execution layer at runtime using optimized environment paths. This completely 
bypasses OS filesystem boundaries, prevents compiled C-extension isolation bugs, and 
drops environment assembly latency down to single-digit milliseconds. 

Featuring parallel multi-threaded downloading, configuration-driven script tracking, 
and static AST-driven codebase tidying, pim delivers a unified, zero-overhead workflow 
for Python developers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pim.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
}


