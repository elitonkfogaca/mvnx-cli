package maven

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// MavenCentralSearchURL is the base URL for Maven Central search API
	MavenCentralSearchURL = "https://search.maven.org/solrsearch/select"

	// DefaultTimeout for HTTP requests
	DefaultTimeout = 10 * time.Second
)

// Client is an HTTP client for Maven Central API.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new Maven Central API client.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		baseURL: MavenCentralSearchURL,
	}
}

// SearchResponse represents the response from Maven Central search API.
type SearchResponse struct {
	Response struct {
		NumFound int        `json:"numFound"`
		Docs     []Document `json:"docs"`
	} `json:"response"`
}

// Document represents a single artifact in the search results.
type Document struct {
	GroupID       string `json:"g"`
	ArtifactID    string `json:"a"`
	LatestVersion string `json:"latestVersion"`
	// Score is not directly in JSON, will be computed
}

// Search performs a search query against Maven Central.
func (c *Client) Search(query string, rows int) (*SearchResponse, error) {
	params := url.Values{}
	params.Add("q", query)
	params.Add("rows", fmt.Sprintf("%d", rows))
	params.Add("wt", "json")

	fullURL := fmt.Sprintf("%s?%s", c.baseURL, params.Encode())

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to query Maven Central: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Maven Central API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &searchResp, nil
}

// SearchByCoordinates performs an exact search by groupId and artifactId.
func (c *Client) SearchByCoordinates(groupID, artifactID string) (*SearchResponse, error) {
	query := fmt.Sprintf("g:\"%s\" AND a:\"%s\"", groupID, artifactID)
	return c.Search(query, 1)
}

// IsStableVersion checks if a version string represents a stable release.
// Excludes SNAPSHOT, alpha, beta, RC, M (milestone) versions.
func IsStableVersion(version string) bool {
	lower := strings.ToLower(version)

	unstableMarkers := []string{
		"-snapshot",
		"-alpha",
		"-beta",
		"-rc",
		"-m",
		".alpha",
		".beta",
		".rc",
		".m",
	}

	for _, marker := range unstableMarkers {
		if strings.Contains(lower, marker) {
			return false
		}
	}

	return true
}
