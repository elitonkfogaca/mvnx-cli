package domain

// PomRepository defines the interface for pom.xml operations.
type PomRepository interface {
	// Load reads and parses the pom.xml file.
	Load(path string) error
	
	// AddDependency adds or updates a dependency in the pom.xml.
	// If the dependency already exists (same groupId:artifactId), it updates the version.
	AddDependency(dep *Dependency) error
	
	// RemoveDependency removes a dependency by artifactId.
	RemoveDependency(artifactID string) error
	
	// HasDependency checks if a dependency exists by groupId and artifactId.
	HasDependency(groupID, artifactID string) bool
	
	// Save writes the pom.xml back to disk, preserving formatting.
	Save() error
	
	// GetDependencies returns all dependencies in the pom.xml.
	GetDependencies() ([]*Dependency, error)
}
