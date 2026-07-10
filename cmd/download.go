/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download and cache a portable standalone CPython runtime version.",
	Long: `Streams a pre-compiled, architecture-optimized CPython tarball from the remote index
based on your current OS and CPU platform. It unpacks the isolated environment directly
into the global pim runtime manager directory for instant workspace availability.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented %q yet\n", cmd.Short)
	},
}

func init() {
	pyCmd.AddCommand(downloadCmd)
}
