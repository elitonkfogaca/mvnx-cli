# Contributing to mvnx

Thank you for your interest in contributing to mvnx!

## Principles

When contributing to mvnx, please keep these principles in mind:

- **Keep Maven compatibility** - mvnx generates standard Maven projects
- **Avoid magic behavior** - All operations should be transparent and predictable
- **Maintain deterministic outputs** - Same input should always produce same output
- **Follow Go idioms** - Write idiomatic Go code

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/elitonkfogaca/mvnx-cli.git
cd mvnx-cli
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
make build
```

4. Run tests:
```bash
make test
```

## Project Structure

The project follows Clean Architecture:

```
mvnx-cli/
├── cmd/mvnx/              # CLI entry point
├── internal/
│   ├── cli/               # CLI layer (Cobra commands)
│   ├── app/               # Application layer (use cases)
│   ├── domain/            # Domain layer (models & interfaces)
│   └── infrastructure/    # Infrastructure layer (implementations)
│       ├── maven/         # Maven Central API client
│       ├── xml/           # POM XML manipulation
│       └── fs/            # File system operations
└── test/integration/      # Integration tests
```

## Code Style

- Follow standard Go formatting (`go fmt`)
- Keep functions small and focused
- Separate domain logic from infrastructure
- Write tests for new functionality
- Add godoc comments for exported types and functions

## Testing

- Write unit tests for all new functionality
- Tests should be in `*_test.go` files in the same package
- Use `github.com/stretchr/testify` for assertions
- Aim for >80% code coverage

Run tests:
```bash
make test
```

Run tests with coverage:
```bash
make test-coverage
```

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and ensure they pass (`make test`)
5. Format your code (`make fmt`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to your fork (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Pull Request Requirements

- All tests must pass
- Code must be formatted with `go fmt`
- New functionality must include tests
- Update documentation if behavior changes
- Describe what your PR does and why

## Reporting Issues

When reporting issues, please include:

- mvnx version (`mvnx --version`)
- Go version (`go version`)
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Relevant pom.xml content (if applicable)

## Questions?

Feel free to open an issue for questions or discussions about contributing.

Thank you for contributing to mvnx!
