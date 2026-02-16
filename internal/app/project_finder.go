package app

import (
	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
	"github.com/elitonkfogaca/mvnx-cli/internal/infrastructure/fs"
)

// ProjectFinder helps locate Maven projects.
type ProjectFinder struct {
	initializer *fs.ProjectInitializer
}

// NewProjectFinder creates a new ProjectFinder.
func NewProjectFinder() *ProjectFinder {
	return &ProjectFinder{
		initializer: fs.NewProjectInitializer(),
	}
}

// FindProject searches for a Maven project starting from the given path.
// It walks up the directory tree looking for pom.xml.
func (pf *ProjectFinder) FindProject(startPath string) (*domain.Project, error) {
	pomPath, err := pf.initializer.FindPomXML(startPath)
	if err != nil {
		return nil, err
	}
	
	// Get the directory containing the pom.xml
	projectPath := pomPath[:len(pomPath)-len("/pom.xml")]
	
	return domain.NewProject(projectPath, pomPath), nil
}
