package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDependency(t *testing.T) {
	tests := []struct {
		name        string
		groupID     string
		artifactID  string
		version     string
		scope       string
		wantErr     bool
		expectedScope string
	}{
		{
			name:        "valid dependency with compile scope",
			groupID:     "org.springframework",
			artifactID:  "spring-core",
			version:     "5.3.0",
			scope:       "compile",
			wantErr:     false,
			expectedScope: "compile",
		},
		{
			name:        "valid dependency with default scope",
			groupID:     "junit",
			artifactID:  "junit",
			version:     "4.13.2",
			scope:       "",
			wantErr:     false,
			expectedScope: "compile",
		},
		{
			name:        "valid dependency with test scope",
			groupID:     "org.mockito",
			artifactID:  "mockito-core",
			version:     "3.11.2",
			scope:       "test",
			wantErr:     false,
			expectedScope: "test",
		},
		{
			name:        "invalid scope",
			groupID:     "org.example",
			artifactID:  "example",
			version:     "1.0.0",
			scope:       "invalid",
			wantErr:     true,
		},
		{
			name:        "empty groupID",
			groupID:     "",
			artifactID:  "example",
			version:     "1.0.0",
			scope:       "compile",
			wantErr:     true,
		},
		{
			name:        "empty artifactID",
			groupID:     "org.example",
			artifactID:  "",
			version:     "1.0.0",
			scope:       "compile",
			wantErr:     true,
		},
		{
			name:        "empty version",
			groupID:     "org.example",
			artifactID:  "example",
			version:     "",
			scope:       "compile",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dep, err := NewDependency(tt.groupID, tt.artifactID, tt.version, tt.scope)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, dep)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, dep)
				assert.Equal(t, tt.groupID, dep.GroupID)
				assert.Equal(t, tt.artifactID, dep.ArtifactID)
				assert.Equal(t, tt.version, dep.Version)
				assert.Equal(t, tt.expectedScope, dep.Scope)
			}
		})
	}
}

func TestDependency_String(t *testing.T) {
	tests := []struct {
		name       string
		dependency *Dependency
		expected   string
	}{
		{
			name: "compile scope (default)",
			dependency: &Dependency{
				GroupID:    "org.example",
				ArtifactID: "my-lib",
				Version:    "1.0.0",
				Scope:      "compile",
			},
			expected: "org.example:my-lib:1.0.0",
		},
		{
			name: "test scope",
			dependency: &Dependency{
				GroupID:    "junit",
				ArtifactID: "junit",
				Version:    "4.13.2",
				Scope:      "test",
			},
			expected: "junit:junit:4.13.2 (scope: test)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dependency.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDependency_Coordinates(t *testing.T) {
	dep := &Dependency{
		GroupID:    "org.example",
		ArtifactID: "my-lib",
		Version:    "1.0.0",
		Scope:      "compile",
	}
	
	assert.Equal(t, "org.example:my-lib", dep.Coordinates())
}
