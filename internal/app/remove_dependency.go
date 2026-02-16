package app

import (
	"fmt"

	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
)

// RemoveDependencyService handles removing dependencies from a project.
type RemoveDependencyService struct {
	pomRepository domain.PomRepository
}

// NewRemoveDependencyService creates a new RemoveDependencyService.
func NewRemoveDependencyService(pomRepository domain.PomRepository) *RemoveDependencyService {
	return &RemoveDependencyService{
		pomRepository: pomRepository,
	}
}

// Remove removes a dependency from the pom.xml by artifactId.
func (s *RemoveDependencyService) Remove(artifactID string) error {
	if err := s.pomRepository.RemoveDependency(artifactID); err != nil {
		return fmt.Errorf("failed to remove dependency: %w", err)
	}
	
	if err := s.pomRepository.Save(); err != nil {
		return fmt.Errorf("failed to save pom.xml: %w", err)
	}
	
	return nil
}

// LoadPom loads the pom.xml from the specified path.
func (s *RemoveDependencyService) LoadPom(path string) error {
	return s.pomRepository.Load(path)
}
