package app

import (
	"fmt"

	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
)

// AddDependencyService handles adding dependencies to a project.
type AddDependencyService struct {
	resolver      domain.Resolver
	pomRepository domain.PomRepository
}

// NewAddDependencyService creates a new AddDependencyService.
func NewAddDependencyService(resolver domain.Resolver, pomRepository domain.PomRepository) *AddDependencyService {
	return &AddDependencyService{
		resolver:      resolver,
		pomRepository: pomRepository,
	}
}

// SearchResult represents the result of a dependency search that may need user selection.
type SearchResult struct {
	// Results is the list of artifacts found
	Results []*domain.ArtifactSearchResult
	
	// NeedsSelection indicates if user needs to select from multiple results
	NeedsSelection bool
}

// Search searches for dependencies matching the query.
// Returns a SearchResult that may require user selection if multiple artifacts are found.
func (s *AddDependencyService) Search(query string) (*SearchResult, error) {
	results, err := s.resolver.Resolve(query)
	if err != nil {
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, fmt.Errorf("no artifacts found for query: %s", query)
	}
	
	// If there's only one result or it was an exact match, no selection needed
	needsSelection := len(results) > 1
	
	return &SearchResult{
		Results:        results,
		NeedsSelection: needsSelection,
	}, nil
}

// Add adds a dependency to the pom.xml.
// The artifact parameter should be an ArtifactSearchResult (from Search).
// The scope parameter specifies the dependency scope (compile, test, provided, runtime).
func (s *AddDependencyService) Add(artifact *domain.ArtifactSearchResult, scope string) error {
	// Convert artifact to dependency
	dep, err := artifact.ToDependency(scope)
	if err != nil {
		return err
	}
	
	// Check if dependency already exists
	if s.pomRepository.HasDependency(dep.GroupID, dep.ArtifactID) {
		// Update existing dependency (silent update)
		if err := s.pomRepository.AddDependency(dep); err != nil {
			return fmt.Errorf("failed to update dependency: %w", err)
		}
	} else {
		// Add new dependency
		if err := s.pomRepository.AddDependency(dep); err != nil {
			return fmt.Errorf("failed to add dependency: %w", err)
		}
	}
	
	// Save the pom.xml
	if err := s.pomRepository.Save(); err != nil {
		return fmt.Errorf("failed to save pom.xml: %w", err)
	}
	
	return nil
}

// LoadPom loads the pom.xml from the specified path.
func (s *AddDependencyService) LoadPom(path string) error {
	return s.pomRepository.Load(path)
}
