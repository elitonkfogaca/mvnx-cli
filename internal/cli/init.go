package cli

import (
	"fmt"
	"os"

	"github.com/elitonkfogaca/mvnx-cli/internal/app"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Maven project",
	Long: `Initialize a new Maven project in the current directory.
Creates a minimal pom.xml and standard directory structure.`,
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	
	// Initialize project
	service := app.NewInitProjectService()
	if err := service.Init(cwd); err != nil {
		return err
	}
	
	fmt.Println("✓ Initialized Maven project")
	fmt.Println("  Created pom.xml")
	fmt.Println("  Created src/main/java")
	fmt.Println("  Created src/test/java")
	
	return nil
}
