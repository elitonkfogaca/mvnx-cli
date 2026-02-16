package xml

import (
	"fmt"

	"github.com/beevik/etree"

	"github.com/elitonkfogaca/mvnx-cli/internal/domain"
)

// PomRepository implements the domain.PomRepository interface using etree for XML manipulation.
type PomRepository struct {
	doc      *etree.Document
	filePath string
}

// NewPomRepository creates a new PomRepository instance.
func NewPomRepository() *PomRepository {
	return &PomRepository{}
}

// Load reads and parses the pom.xml file.
func (p *PomRepository) Load(path string) error {
	doc := etree.NewDocument()

	if err := doc.ReadFromFile(path); err != nil {
		return fmt.Errorf("failed to read pom.xml: %w", err)
	}

	p.doc = doc
	p.filePath = path

	return nil
}

// AddDependency adds or updates a dependency in the pom.xml.
func (p *PomRepository) AddDependency(dep *domain.Dependency) error {
	if p.doc == nil {
		return fmt.Errorf("no pom.xml loaded")
	}

	root := p.doc.Root()
	if root == nil {
		return fmt.Errorf("invalid pom.xml: no root element")
	}

	// Find or create <dependencies> element
	dependencies := root.SelectElement("dependencies")
	if dependencies == nil {
		dependencies = root.CreateElement("dependencies")
	}

	// Check if dependency already exists
	existingDep := p.findDependency(dependencies, dep.GroupID, dep.ArtifactID)

	if existingDep != nil {
		// Update existing dependency
		p.updateDependencyElement(existingDep, dep)
	} else {
		// Add new dependency
		p.createDependencyElement(dependencies, dep)
	}

	return nil
}

// RemoveDependency removes a dependency by artifactId.
func (p *PomRepository) RemoveDependency(artifactID string) error {
	if p.doc == nil {
		return fmt.Errorf("no pom.xml loaded")
	}

	root := p.doc.Root()
	dependencies := root.SelectElement("dependencies")
	if dependencies == nil {
		return fmt.Errorf("dependency not found: %s", artifactID)
	}

	// Find all dependency elements
	for _, dep := range dependencies.SelectElements("dependency") {
		artifactElem := dep.SelectElement("artifactId")
		if artifactElem != nil && artifactElem.Text() == artifactID {
			dependencies.RemoveChild(dep)
			return nil
		}
	}

	return fmt.Errorf("dependency not found: %s", artifactID)
}

// HasDependency checks if a dependency exists by groupId and artifactId.
func (p *PomRepository) HasDependency(groupID, artifactID string) bool {
	if p.doc == nil {
		return false
	}

	root := p.doc.Root()
	dependencies := root.SelectElement("dependencies")
	if dependencies == nil {
		return false
	}

	return p.findDependency(dependencies, groupID, artifactID) != nil
}

// GetDependencies returns all dependencies in the pom.xml.
func (p *PomRepository) GetDependencies() ([]*domain.Dependency, error) {
	if p.doc == nil {
		return nil, fmt.Errorf("no pom.xml loaded")
	}

	root := p.doc.Root()
	dependencies := root.SelectElement("dependencies")
	if dependencies == nil {
		return []*domain.Dependency{}, nil
	}

	var result []*domain.Dependency

	for _, dep := range dependencies.SelectElements("dependency") {
		groupElem := dep.SelectElement("groupId")
		artifactElem := dep.SelectElement("artifactId")
		versionElem := dep.SelectElement("version")
		scopeElem := dep.SelectElement("scope")

		if groupElem == nil || artifactElem == nil || versionElem == nil {
			continue
		}

		scope := "compile"
		if scopeElem != nil && scopeElem.Text() != "" {
			scope = scopeElem.Text()
		}

		dependency, err := domain.NewDependency(
			groupElem.Text(),
			artifactElem.Text(),
			versionElem.Text(),
			scope,
		)
		if err != nil {
			continue
		}

		result = append(result, dependency)
	}

	return result, nil
}

// Save writes the pom.xml back to disk, preserving formatting.
func (p *PomRepository) Save() error {
	if p.doc == nil {
		return fmt.Errorf("no pom.xml loaded")
	}

	p.doc.Indent(2)

	if err := p.doc.WriteToFile(p.filePath); err != nil {
		return fmt.Errorf("failed to write pom.xml: %w", err)
	}

	return nil
}

// findDependency finds a dependency element by groupId and artifactId.
func (p *PomRepository) findDependency(dependencies *etree.Element, groupID, artifactID string) *etree.Element {
	for _, dep := range dependencies.SelectElements("dependency") {
		groupElem := dep.SelectElement("groupId")
		artifactElem := dep.SelectElement("artifactId")

		if groupElem != nil && artifactElem != nil &&
			groupElem.Text() == groupID && artifactElem.Text() == artifactID {
			return dep
		}
	}
	return nil
}

// updateDependencyElement updates an existing dependency element.
func (p *PomRepository) updateDependencyElement(elem *etree.Element, dep *domain.Dependency) {
	versionElem := elem.SelectElement("version")
	if versionElem != nil {
		versionElem.SetText(dep.Version)
	}

	// Update or add scope if not compile
	scopeElem := elem.SelectElement("scope")
	if dep.Scope != "compile" {
		if scopeElem == nil {
			scopeElem = elem.CreateElement("scope")
		}
		scopeElem.SetText(dep.Scope)
	} else if scopeElem != nil {
		// Remove scope element if it's compile (default)
		elem.RemoveChild(scopeElem)
	}
}

// createDependencyElement creates a new dependency element.
func (p *PomRepository) createDependencyElement(dependencies *etree.Element, dep *domain.Dependency) {
	depElem := dependencies.CreateElement("dependency")

	groupElem := depElem.CreateElement("groupId")
	groupElem.SetText(dep.GroupID)

	artifactElem := depElem.CreateElement("artifactId")
	artifactElem.SetText(dep.ArtifactID)

	versionElem := depElem.CreateElement("version")
	versionElem.SetText(dep.Version)

	// Only add scope if not compile (default)
	if dep.Scope != "compile" {
		scopeElem := depElem.CreateElement("scope")
		scopeElem.SetText(dep.Scope)
	}
}
