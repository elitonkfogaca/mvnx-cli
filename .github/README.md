# GitHub Workflows

This directory contains GitHub Actions workflows for CI/CD.

## Workflows

### CI (`ci.yml`)
Runs on every push and pull request:
- Tests on Linux, macOS, and Windows
- Linting with golangci-lint
- Coverage reporting

### Release (`release.yml`)
Runs on version tags (e.g., `v1.0.0`):
- Builds binaries for all platforms
- Creates GitHub release
- Uploads artifacts

## Creating a Release

1. Update VERSION file
2. Commit changes
3. Create and push a tag:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```
4. GitHub Actions will automatically build and publish the release
