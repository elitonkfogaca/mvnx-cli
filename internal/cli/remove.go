package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/elitonkfogaca/mvnx-cli/internal/app"
	"github.com/elitonkfogaca/mvnx-cli/internal/infrastructure/xml"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <artifactId>",
	Short: "Remove a dependency from the project",
	Long:  `Remove a dependency from the project's pom.xml by its artifactId.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runRemove,
}

func runRemove(cmd *cobra.Command, args []string) error {
	artifactID := args[0]

	// Find project
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	projectFinder := app.NewProjectFinder()
	project, err := projectFinder.FindProject(cwd)
	if err != nil {
		return fmt.Errorf("no Maven project found")
	}

	if verbose {
		fmt.Printf("Found pom.xml at: %s\n", project.PomLocation)
	}

	// Create service
	pomRepo := xml.NewPomRepository()
	service := app.NewRemoveDependencyService(pomRepo)

	// Load pom.xml
	if err := service.LoadPom(project.PomLocation); err != nil {
		return fmt.Errorf("failed to load pom.xml: %w", err)
	}

	// Remove dependency
	if err := service.Remove(artifactID); err != nil {
		return err
	}

	fmt.Printf("✓ Removed %s\n", artifactID)

	return nil
}
