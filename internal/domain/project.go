package domain

// Project represents a Maven project in the file system.
type Project struct {
	// Path is the directory containing the pom.xml
	Path string

	// PomLocation is the full path to the pom.xml file
	PomLocation string
}

// NewProject creates a new Project instance.
func NewProject(path, pomLocation string) *Project {
	return &Project{
		Path:        path,
		PomLocation: pomLocation,
	}
}
