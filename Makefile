.PHONY: build test clean install help

# Build the binary
build:
	@echo "Building mvnx..."
	@go build -o mvnx ./cmd/mvnx
	@echo "✓ Build complete: ./mvnx"

# Run tests
test:
	@echo "Running tests..."
	@go test ./... -v

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f mvnx
	@rm -f coverage.out coverage.html
	@echo "✓ Clean complete"

# Install binary to GOBIN
install:
	@echo "Installing mvnx..."
	@go install ./cmd/mvnx
	@echo "✓ Installed to GOBIN"

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Format complete"

# Run all checks (fmt, lint, test)
check: fmt lint test
	@echo "✓ All checks passed"

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the mvnx binary"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Remove build artifacts"
	@echo "  install       - Install binary to GOBIN"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  check         - Run fmt, lint, and test"
	@echo "  help          - Show this help message"
