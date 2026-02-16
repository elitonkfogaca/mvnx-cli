package maven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStableVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{
			name:    "stable version - semantic",
			version: "1.2.3",
			want:    true,
		},
		{
			name:    "stable version - with patch",
			version: "5.3.15",
			want:    true,
		},
		{
			name:    "SNAPSHOT version",
			version: "1.0.0-SNAPSHOT",
			want:    false,
		},
		{
			name:    "alpha version",
			version: "2.0.0-alpha",
			want:    false,
		},
		{
			name:    "beta version",
			version: "2.0.0-beta.1",
			want:    false,
		},
		{
			name:    "RC version",
			version: "3.0.0-RC1",
			want:    false,
		},
		{
			name:    "milestone version",
			version: "4.0.0-M2",
			want:    false,
		},
		{
			name:    "alpha with dot",
			version: "1.0.0.alpha",
			want:    false,
		},
		{
			name:    "SNAPSHOT uppercase",
			version: "1.0.0-SNAPSHOT",
			want:    false,
		},
		{
			name:    "snapshot lowercase",
			version: "1.0.0-snapshot",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStableVersion(tt.version)
			assert.Equal(t, tt.want, result)
		})
	}
}
