package app

import (
	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
)

// SearchArtifactsService handles searching for artifacts.
type SearchArtifactsService struct {
	resolver domain.Resolver
}

// NewSearchArtifactsService creates a new SearchArtifactsService.
func NewSearchArtifactsService(resolver domain.Resolver) *SearchArtifactsService {
	return &SearchArtifactsService{
		resolver: resolver,
	}
}

// Search searches for artifacts matching the query.
// Returns up to 10 results ordered by relevance.
func (s *SearchArtifactsService) Search(query string) ([]*domain.ArtifactSearchResult, error) {
	results, err := s.resolver.Resolve(query)
	if err != nil {
		return nil, err
	}
	
	// Limit to top 5 results for display
	if len(results) > 5 {
		results = results[:5]
	}
	
	return results, nil
}
