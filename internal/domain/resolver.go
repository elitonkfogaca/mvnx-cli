package domain

// Resolver defines the interface for resolving Maven artifacts.
type Resolver interface {
	// Resolve searches for artifacts matching the query.
	// Returns a list of results ordered by relevance score.
	Resolve(query string) ([]*ArtifactSearchResult, error)

	// ResolveExact performs an exact lookup for a specific groupId:artifactId.
	ResolveExact(groupID, artifactID string) (*ArtifactSearchResult, error)
}
