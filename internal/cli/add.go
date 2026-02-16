package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/elitonkfogaca/mvnx-cli/internal/app"
	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
	"github.com/elitonkfogaca/mvnx-cli/internal/infrastructure/maven"
	"github.com/elitonkfogaca/mvnx-cli/internal/infrastructure/xml"
	"github.com/spf13/cobra"
)

var (
	// scope flag for add command
	scope string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <query>",
	Short: "Add a dependency to the project",
	Long: `Add a dependency to the project's pom.xml.
Query can be a simple search term (e.g., "lombok") or an exact coordinate (e.g., "org.projectlombok:lombok").`,
	Args: cobra.ExactArgs(1),
	RunE: runAdd,
}

func init() {
	addCmd.Flags().StringVar(&scope, "scope", "compile", "dependency scope (compile, test, provided, runtime)")
}

func runAdd(cmd *cobra.Command, args []string) error {
	query := args[0]
	
	// Validate scope
	validScopes := map[string]bool{
		"compile":  true,
		"test":     true,
		"provided": true,
		"runtime":  true,
	}
	if !validScopes[scope] {
		return fmt.Errorf("invalid scope: %s (valid: compile, test, provided, runtime)", scope)
	}
	
	// Find project
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	
	projectFinder := app.NewProjectFinder()
	project, err := projectFinder.FindProject(cwd)
	if err != nil {
		return fmt.Errorf("no Maven project found. Run 'mvnx init' to create one.")
	}
	
	if verbose {
		fmt.Printf("Found pom.xml at: %s\n", project.PomLocation)
	}
	
	// Create services
	resolver := maven.NewResolver()
	pomRepo := xml.NewPomRepository()
	service := app.NewAddDependencyService(resolver, pomRepo)
	
	// Load pom.xml
	if err := service.LoadPom(project.PomLocation); err != nil {
		return fmt.Errorf("failed to load pom.xml: %w", err)
	}
	
	// Search for artifacts
	if verbose {
		fmt.Printf("Searching for: %s\n", query)
	}
	
	searchResult, err := service.Search(query)
	if err != nil {
		return err
	}
	
	// Select artifact
	var selectedArtifact *domain.ArtifactSearchResult
	
	if searchResult.NeedsSelection {
		// Interactive selection
		selectedArtifact, err = selectArtifact(searchResult.Results)
		if err != nil {
			return err
		}
	} else {
		// Use the only result
		selectedArtifact = searchResult.Results[0]
	}
	
	// Add the dependency
	if err := service.Add(selectedArtifact, scope); err != nil {
		return err
	}
	
	fmt.Printf("✓ Added %s\n", selectedArtifact.String())
	
	return nil
}

// selectArtifact presents an interactive selection menu and returns the chosen artifact
func selectArtifact(results []*domain.ArtifactSearchResult) (*domain.ArtifactSearchResult, error) {
	fmt.Println("\nMultiple artifacts found:")
	fmt.Println()
	
	for i, result := range results {
		fmt.Printf("[%d] %s:%s\n", i+1, result.GroupID, result.ArtifactID)
		fmt.Printf("    Version: %s\n", result.LatestVersion)
		fmt.Println()
	}
	
	fmt.Printf("Select artifact (1-%d): ", len(results))
	
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}
	
	input = strings.TrimSpace(input)
	selection, err := strconv.Atoi(input)
	if err != nil || selection < 1 || selection > len(results) {
		return nil, fmt.Errorf("invalid selection: %s", input)
	}
	
	return results[selection-1], nil
}
