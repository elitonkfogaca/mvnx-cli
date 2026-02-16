package cli

import (
	"github.com/spf13/cobra"
)

var (
	// Verbose flag for detailed output
	verbose bool

	// Version information
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// SetVersion sets the version information for the CLI
func SetVersion(v, c, d string) {
	version = v
	commit = c
	date = d
}

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "mvnx",
	Short: "Modern Dependency Experience for Maven",
	Long: `mvnx is a CLI tool that enhances the Maven developer experience.
It provides a modern, intelligent layer on top of standard Maven projects.`,
	Version: version,
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Customize version template
	rootCmd.SetVersionTemplate(`{{printf "mvnx version %s\n" .Version}}{{printf "commit: %s\n" (index .Annotations "commit")}}{{printf "built: %s\n" (index .Annotations "date")}}`)
	rootCmd.Annotations = map[string]string{
		"commit": commit,
		"date":   date,
	}

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(searchCmd)
}
