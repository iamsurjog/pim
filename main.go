/*
Copyright © 2026 Sujatro Ganguli iamsurjog@gmail.com
*/
package main

import (
	"pim/cmd"

	"fmt"
	"os"
	"runtime"
)

func main() {

	// TODO: Implement for other OSes as well
	osType := runtime.GOOS

	// Example conditional check
	if osType == "windows" {
		fmt.Println("Not built for windows yet. Only supports linux")
		os.Exit(0)
	} else if osType == "linux" {
		fmt.Println("Running on Linux.")
	} else if osType == "darwin" {
		fmt.Println("Not built for macOS yet. Only supports linux")
		os.Exit(0)
	}
	cmd.Execute()
}
