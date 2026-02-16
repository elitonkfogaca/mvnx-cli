package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/elitonkfogaca/mvnx-cli/internal/app"
	"github.com/elitonkfogaca/mvnx-cli/internal/infrastructure/maven"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for artifacts in Maven Central",
	Long:  `Search Maven Central for artifacts matching the query. Shows top 5 results.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runSearch,
}

func runSearch(cmd *cobra.Command, args []string) error {
	query := args[0]

	// Create service
	resolver := maven.NewResolver()
	service := app.NewSearchArtifactsService(resolver)

	if verbose {
		fmt.Printf("Searching Maven Central for: %s\n\n", query)
	}

	// Search
	results, err := service.Search(query)
	if err != nil {
		return err
	}

	// Display results
	if len(results) == 0 {
		fmt.Println("No artifacts found")
		return nil
	}

	// Calculate column widths for better formatting
	maxGroupLen := 0
	maxArtifactLen := 0
	for _, result := range results {
		if len(result.GroupID) > maxGroupLen {
			maxGroupLen = len(result.GroupID)
		}
		if len(result.ArtifactID) > maxArtifactLen {
			maxArtifactLen = len(result.ArtifactID)
		}
	}

	// Print header
	fmt.Printf("%-*s  %-*s  %s\n", maxGroupLen, "GROUP ID", maxArtifactLen, "ARTIFACT ID", "LATEST VERSION")
	fmt.Println(strings.Repeat("-", maxGroupLen+maxArtifactLen+30))

	// Print results
	for _, result := range results {
		fmt.Printf("%-*s  %-*s  %s\n",
			maxGroupLen, result.GroupID,
			maxArtifactLen, result.ArtifactID,
			result.LatestVersion)
	}

	return nil
}
