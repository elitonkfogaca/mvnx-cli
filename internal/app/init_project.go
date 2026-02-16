package app

import (
	"github.com/elitonkfogaca/mvnx-cli/internal/infrastructure/fs"
)

// InitProjectService handles project initialization.
type InitProjectService struct {
	initializer *fs.ProjectInitializer
}

// NewInitProjectService creates a new InitProjectService.
func NewInitProjectService() *InitProjectService {
	return &InitProjectService{
		initializer: fs.NewProjectInitializer(),
	}
}

// Init initializes a new Maven project in the specified directory.
func (s *InitProjectService) Init(path string) error {
	return s.initializer.InitProject(path)
}
