package domain

import "fmt"

// ArtifactSearchResult represents a search result from Maven Central.
type ArtifactSearchResult struct {
	GroupID       string
	ArtifactID    string
	LatestVersion string
	Score         float64
}

// NewArtifactSearchResult creates a new ArtifactSearchResult.
func NewArtifactSearchResult(groupID, artifactID, latestVersion string, score float64) *ArtifactSearchResult {
	return &ArtifactSearchResult{
		GroupID:       groupID,
		ArtifactID:    artifactID,
		LatestVersion: latestVersion,
		Score:         score,
	}
}

// String returns a formatted string representation of the artifact.
func (a *ArtifactSearchResult) String() string {
	return fmt.Sprintf("%s:%s:%s", a.GroupID, a.ArtifactID, a.LatestVersion)
}

// Coordinates returns the Maven coordinates without version.
func (a *ArtifactSearchResult) Coordinates() string {
	return fmt.Sprintf("%s:%s", a.GroupID, a.ArtifactID)
}

// ToDependency converts the search result to a Dependency with the specified scope.
func (a *ArtifactSearchResult) ToDependency(scope string) (*Dependency, error) {
	return NewDependency(a.GroupID, a.ArtifactID, a.LatestVersion, scope)
}
