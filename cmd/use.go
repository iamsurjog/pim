/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch the current project workspace to a specific Python version.",
	Long: `Mutates the 'python = "<version>"' key inside the local project configuration file
and points the path routing variables to the requested binary directory. This re-routes
execution instantly without altering system-wide environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	pyCmd.AddCommand(useCmd)
}
