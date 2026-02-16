# mvnx

[![CI](https://github.com/elitonkfogaca/mvnx-cli/workflows/CI/badge.svg)](https://github.com/elitonkfogaca/mvnx-cli/actions)
[![Release](https://img.shields.io/github/v/release/elitonkfogaca/mvnx-cli)](https://github.com/elitonkfogaca/mvnx-cli/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)

Modern Dependency Experience for Maven.

mvnx is a lightweight, cross-platform CLI tool that enhances Maven developer experience.

It provides a modern, intelligent interface for dependency management on top of standard Maven projects — without replacing Maven and without introducing lock-in.

**Platforms:** Linux, macOS, Windows (amd64, arm64)

---

## Why mvnx?

Managing dependencies in Maven still requires manual XML editing and browser searches.

mvnx modernizes that workflow with a simple CLI experience.

**Instead of:**

1. Searching artifacts in a browser  
2. Copying groupId and artifactId  
3. Checking versions  
4. Editing XML manually  

**Just run:**

```bash
mvnx add spring-boot-starter-web
```

mvnx handles the search, version resolution, and XML updates automatically.

---

## Features (v1)

- `mvnx init` — Initialize a minimal Maven project
- `mvnx add <query>` — Add dependency with automatic version resolution
- `mvnx remove <artifactId>` — Remove dependency
- `mvnx search <query>` — Search Maven Central

---

## Design Principles

- Zero lock-in
- 100% Maven compatible
- Deterministic behavior
- CLI-first experience
- Intelligent dependency resolution

---

## Example

```
mvnx init
mvnx add lombok
mvnx add spring-web
mvnx search postgres
```  

---

## System Requirements

- **Operating System:** Linux, macOS, or Windows
- **Maven:** Not required for mvnx itself (operates on pom.xml files)
- **Java:** Required only for building/running your Maven projects
- **Network:** Internet connection for Maven Central queries

mvnx is a standalone binary with no runtime dependencies.

---

## Releases

mvnx follows [semantic versioning](https://semver.org/).

Pre-built binaries are available for:
- **Linux:** amd64, arm64
- **macOS:** amd64 (Intel), arm64 (Apple Silicon)
- **Windows:** amd64

Download the latest release from the [releases page](https://github.com/elitonkfogaca/mvnx-cli/releases).

---

## Roadmap

### v1.0.0
- init
- add
- remove
- search

### v1.1.0
- upgrade
- outdated

### v2.0.0
- lockfile (mvnx.lock)
- dependency graph
- doctor command

### v3.0.0
- Gradle support
- Multi-module support
- Plugin system

---

## Installation

mvnx is distributed as a single binary with no dependencies. Download the appropriate version for your platform and add it to your PATH.

> **📖 For detailed installation instructions, see [INSTALL.md](INSTALL.md)**

### Quick Install (macOS/Linux)

```bash
curl -fsSL https://raw.githubusercontent.com/elitonkfogaca/mvnx-cli/main/install.sh | bash
```

This script automatically detects your OS and architecture, downloads the latest release, and installs it to `/usr/local/bin`.

### macOS

**Via Homebrew (coming soon):**
```bash
brew tap elitonkfogaca/tap
brew install mvnx
```

**Manual installation:**
```bash
# For Apple Silicon (M1/M2/M3)
curl -L https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_darwin_arm64.tar.gz | tar xz
sudo mv mvnx /usr/local/bin/

# For Intel Macs
curl -L https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_darwin_amd64.tar.gz | tar xz
sudo mv mvnx /usr/local/bin/
```

### Linux

```bash
# For x86_64
curl -L https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_linux_amd64.tar.gz | tar xz
sudo mv mvnx /usr/local/bin/

# For ARM64
curl -L https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_linux_arm64.tar.gz | tar xz
sudo mv mvnx /usr/local/bin/
```

### Windows

**Via PowerShell:**
```powershell
# Download the latest release
Invoke-WebRequest -Uri "https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_windows_amd64.zip" -OutFile "mvnx.zip"

# Extract
Expand-Archive -Path mvnx.zip -DestinationPath .

# Add to PATH (optional - requires admin)
Move-Item mvnx.exe C:\Windows\System32\
```

**Manual:**
1. Download the latest [Windows release](https://github.com/elitonkfogaca/mvnx-cli/releases/latest)
2. Extract `mvnx.exe` from the ZIP file
3. Add the directory containing `mvnx.exe` to your PATH

### From Source

Requires Go 1.21 or later:

```bash
go install github.com/elitonkfogaca/mvnx-cli/cmd/mvnx@latest
```

### Verify Installation

```bash
mvnx --help
mvnx --version
```

### Update

**Homebrew:**
```bash
brew upgrade mvnx
```

**Manual:** Download and install the latest release following the installation instructions above.

**Go install:**
```bash
go install github.com/elitonkfogaca/mvnx-cli/cmd/mvnx@latest
```

### Uninstall

**Homebrew:**
```bash
brew uninstall mvnx
```

**Manual:**
```bash
# macOS/Linux
sudo rm /usr/local/bin/mvnx

# Windows
del C:\Windows\System32\mvnx.exe
```

---

## Quick Start

1. **Create a new Maven project:**
```bash
mkdir my-project && cd my-project
mvnx init
```

2. **Add dependencies:**
```bash
mvnx add spring-boot-starter-web
mvnx add lombok --scope provided
mvnx add junit --scope test
```

3. **Search for artifacts:**
```bash
mvnx search postgresql
```

4. **Remove dependencies:**
```bash
mvnx remove lombok
```

---

## Usage

### `mvnx init`

Initialize a new Maven project in the current directory.

Creates:
- `pom.xml` with Java 17 configuration
- `src/main/java` directory structure
- `src/test/java` directory structure

```bash
mvnx init
```

### `mvnx add <query>`

Add a dependency to your project.

**Examples:**

```bash
# Search by artifact name
mvnx add lombok

# Exact coordinates
mvnx add org.projectlombok:lombok

# With custom scope
mvnx add junit --scope test
mvnx add lombok --scope provided
```

**Interactive Selection:**

When multiple artifacts match your query, mvnx presents an interactive menu:

```
Multiple artifacts found:

[1] org.postgresql:postgresql
    Version: 42.7.0

[2] com.impossibl.pgjdbc-ng:pgjdbc-ng
    Version: 0.8.9

Select artifact (1-2): 1
```

**Scopes:**
- `compile` (default) - Available in all classpaths
- `test` - Only for testing
- `provided` - Expected to be provided by JDK or container
- `runtime` - Not required for compilation, but for execution

### `mvnx search <query>`

Search Maven Central for artifacts.

```bash
mvnx search spring-boot
mvnx search postgresql
```

Shows top 5 results with groupId, artifactId, and latest version.

### `mvnx remove <artifactId>`

Remove a dependency from your project.

```bash
mvnx remove lombok
mvnx remove junit
```

### Verbose Mode

Add `-v` flag for detailed output:

```bash
mvnx add spring-web -v
mvnx search lombok -v
```

---

## How It Works

mvnx works with standard Maven projects:

1. **Zero lock-in:** Generates standard `pom.xml` files
2. **Maven compatible:** Works with existing Maven projects
3. **Preserves formatting:** Maintains your XML formatting and comments
4. **Intelligent resolution:** Finds latest stable versions automatically
5. **Project detection:** Works from any subdirectory in your project

mvnx integrates with Maven Central API to:
- Search for artifacts by name or coordinates
- Resolve latest stable versions (excludes SNAPSHOT, alpha, beta, RC)
- Provide relevance-ranked search results

---

## FAQ

**Q: Does mvnx replace Maven?**  
A: No. mvnx is a CLI tool that enhances Maven by simplifying dependency management. Your projects remain standard Maven projects.

**Q: Can I use mvnx with existing Maven projects?**  
A: Yes! mvnx works with any existing Maven project. Just run mvnx commands in your project directory.

**Q: Does mvnx modify my pom.xml formatting?**  
A: mvnx preserves your XML formatting, comments, and indentation as much as possible.

**Q: Does mvnx work offline?**  
A: Searching and adding dependencies requires internet access to query Maven Central. Other commands work offline.

**Q: Can I use mvnx in CI/CD?**  
A: Yes. mvnx is designed to work in automated environments. It has deterministic behavior and clear exit codes.

**Q: Does mvnx support multi-module projects?**  
A: Not in v1.0. Multi-module support is planned for v3.0 (see Roadmap).

---

## Contributing

We welcome contributions! Whether it's bug reports, feature requests, or code contributions.

**Guidelines:**
- Maintain Maven compatibility
- Keep behavior deterministic
- Follow Go idioms
- Include tests for new features

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

**Development:**
```bash
git clone https://github.com/elitonkfogaca/mvnx-cli.git
cd mvnx-cli
make build
make test
```

---

## Support

- **Issues:** [GitHub Issues](https://github.com/elitonkfogaca/mvnx-cli/issues)
- **Discussions:** [GitHub Discussions](https://github.com/elitonkfogaca/mvnx-cli/discussions)
- **Documentation:** See [docs/](docs/) folder

---

## License

MIT License

Copyright (c) 2026 Eliton Fogaça

See [LICENSE](LICENSE) for details.

---

## Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [etree](https://github.com/beevik/etree) - XML processing
- [Maven Central](https://search.maven.org/) - Artifact repository

Inspired by modern package managers: npm, cargo, uv, and poetry.
