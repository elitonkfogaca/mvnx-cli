package main

import (
	"fmt"
	"os"

	"github.com/elitonkfogaca/mvnx-cli/internal/cli"
)

// Build information. Populated at build-time via ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Set version information for CLI
	cli.SetVersion(version, commit, date)
	
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
