package domain

import "fmt"

// Dependency represents a Maven dependency with its coordinates and scope.
type Dependency struct {
	GroupID    string
	ArtifactID string
	Version    string
	Scope      string // compile, test, provided, runtime
}

// NewDependency creates a new Dependency with validation.
func NewDependency(groupID, artifactID, version, scope string) (*Dependency, error) {
	if groupID == "" {
		return nil, fmt.Errorf("groupID cannot be empty")
	}
	if artifactID == "" {
		return nil, fmt.Errorf("artifactID cannot be empty")
	}
	if version == "" {
		return nil, fmt.Errorf("version cannot be empty")
	}

	// Default scope to compile if not specified
	if scope == "" {
		scope = "compile"
	}

	// Validate scope
	validScopes := map[string]bool{
		"compile":  true,
		"test":     true,
		"provided": true,
		"runtime":  true,
	}
	if !validScopes[scope] {
		return nil, fmt.Errorf("invalid scope: %s (valid: compile, test, provided, runtime)", scope)
	}

	return &Dependency{
		GroupID:    groupID,
		ArtifactID: artifactID,
		Version:    version,
		Scope:      scope,
	}, nil
}

// String returns a formatted string representation of the dependency.
func (d *Dependency) String() string {
	if d.Scope == "compile" || d.Scope == "" {
		return fmt.Sprintf("%s:%s:%s", d.GroupID, d.ArtifactID, d.Version)
	}
	return fmt.Sprintf("%s:%s:%s (scope: %s)", d.GroupID, d.ArtifactID, d.Version, d.Scope)
}

// Coordinates returns the Maven coordinates without version.
func (d *Dependency) Coordinates() string {
	return fmt.Sprintf("%s:%s", d.GroupID, d.ArtifactID)
}
