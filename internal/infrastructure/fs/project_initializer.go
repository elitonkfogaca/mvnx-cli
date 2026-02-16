package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// DefaultPomTemplate is the minimal pom.xml template
	DefaultPomTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0
         http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.example</groupId>
  <artifactId>my-app</artifactId>
  <version>1.0-SNAPSHOT</version>
  <packaging>jar</packaging>

  <name>my-app</name>

  <properties>
    <maven.compiler.source>17</maven.compiler.source>
    <maven.compiler.target>17</maven.compiler.target>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>

  <dependencies>
  </dependencies>

  <build>
    <plugins>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-compiler-plugin</artifactId>
        <version>3.11.0</version>
      </plugin>
    </plugins>
  </build>
</project>
`
)

// ProjectInitializer handles project initialization operations.
type ProjectInitializer struct{}

// NewProjectInitializer creates a new ProjectInitializer instance.
func NewProjectInitializer() *ProjectInitializer {
	return &ProjectInitializer{}
}

// InitProject creates a new Maven project structure in the current directory.
func (pi *ProjectInitializer) InitProject(path string) error {
	// Check if pom.xml already exists
	pomPath := filepath.Join(path, "pom.xml")
	if _, err := os.Stat(pomPath); err == nil {
		return fmt.Errorf("pom.xml already exists in %s", path)
	}

	// Create pom.xml
	if err := os.WriteFile(pomPath, []byte(DefaultPomTemplate), 0644); err != nil {
		return fmt.Errorf("failed to create pom.xml: %w", err)
	}

	// Create directory structure
	dirs := []string{
		filepath.Join(path, "src", "main", "java"),
		filepath.Join(path, "src", "main", "resources"),
		filepath.Join(path, "src", "test", "java"),
		filepath.Join(path, "src", "test", "resources"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

// FindPomXML searches for pom.xml starting from the given directory and walking up the tree.
func (pi *ProjectInitializer) FindPomXML(startPath string) (string, error) {
	currentPath, err := filepath.Abs(startPath)
	if err != nil {
		return "", err
	}

	// Walk up the directory tree
	for {
		pomPath := filepath.Join(currentPath, "pom.xml")

		if _, err := os.Stat(pomPath); err == nil {
			return pomPath, nil
		}

		// Get parent directory
		parentPath := filepath.Dir(currentPath)

		// If we've reached the root, stop
		if parentPath == currentPath {
			break
		}

		currentPath = parentPath
	}

	return "", fmt.Errorf("no pom.xml found in current directory or parent directories")
}
