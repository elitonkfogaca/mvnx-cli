package maven

import (
	"fmt"
	"strings"

	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
)

// Resolver implements the domain.Resolver interface using Maven Central API.
type Resolver struct {
	client *Client
}

// NewResolver creates a new Maven Central resolver.
func NewResolver() *Resolver {
	return &Resolver{
		client: NewClient(),
	}
}

// Resolve searches for artifacts matching the query.
// If query contains ":", it's treated as groupId:artifactId for exact lookup.
// Otherwise, it performs a fuzzy search.
func (r *Resolver) Resolve(query string) ([]*domain.ArtifactSearchResult, error) {
	// Check if query is in groupId:artifactId format
	if strings.Contains(query, ":") {
		parts := strings.SplitN(query, ":", 2)
		if len(parts) == 2 {
			groupID := strings.TrimSpace(parts[0])
			artifactID := strings.TrimSpace(parts[1])
			
			if groupID != "" && artifactID != "" {
				result, err := r.ResolveExact(groupID, artifactID)
				if err != nil {
					return nil, err
				}
				return []*domain.ArtifactSearchResult{result}, nil
			}
		}
	}
	
	// Perform fuzzy search
	return r.fuzzySearch(query)
}

// ResolveExact performs an exact lookup for a specific groupId:artifactId.
func (r *Resolver) ResolveExact(groupID, artifactID string) (*domain.ArtifactSearchResult, error) {
	resp, err := r.client.SearchByCoordinates(groupID, artifactID)
	if err != nil {
		return nil, err
	}
	
	if resp.Response.NumFound == 0 {
		return nil, fmt.Errorf("artifact not found: %s:%s", groupID, artifactID)
	}
	
	doc := resp.Response.Docs[0]
	
	// Ensure version is stable
	version := doc.LatestVersion
	if !IsStableVersion(version) {
		return nil, fmt.Errorf("no stable version found for %s:%s (latest: %s)", groupID, artifactID, version)
	}
	
	return domain.NewArtifactSearchResult(
		doc.GroupID,
		doc.ArtifactID,
		doc.LatestVersion,
		100.0, // Max score for exact match
	), nil
}

// fuzzySearch performs a fuzzy search with multiple results.
func (r *Resolver) fuzzySearch(query string) ([]*domain.ArtifactSearchResult, error) {
	// Search for more results than we'll return to allow filtering
	resp, err := r.client.Search(query, 20)
	if err != nil {
		return nil, err
	}
	
	if resp.Response.NumFound == 0 {
		return nil, fmt.Errorf("no artifacts found for query: %s", query)
	}
	
	results := make([]*domain.ArtifactSearchResult, 0)
	
	for i, doc := range resp.Response.Docs {
		// Filter out unstable versions
		if !IsStableVersion(doc.LatestVersion) {
			continue
		}
		
		// Calculate a simple relevance score based on position
		// First result gets highest score, subsequent results get lower scores
		score := 100.0 - float64(i)*10.0
		if score < 0 {
			score = 0
		}
		
		result := domain.NewArtifactSearchResult(
			doc.GroupID,
			doc.ArtifactID,
			doc.LatestVersion,
			score,
		)
		
		results = append(results, result)
		
		// Stop after we have enough stable results
		if len(results) >= 10 {
			break
		}
	}
	
	if len(results) == 0 {
		return nil, fmt.Errorf("no stable versions found for query: %s", query)
	}
	
	return results, nil
}
