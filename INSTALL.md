# Installation Guide

This guide provides detailed installation instructions for mvnx on different platforms.

## Table of Contents

- [macOS](#macos)
- [Linux](#linux)
- [Windows](#windows)
- [From Source](#from-source)
- [Verification](#verification)
- [Troubleshooting](#troubleshooting)

---

## macOS

### Quick Install (Recommended)

Use the installation script:

```bash
curl -fsSL https://raw.githubusercontent.com/elitonkfogaca/mvnx-cli/main/install.sh | bash
```

### Homebrew (Coming Soon)

```bash
brew tap elitonkfogaca/tap
brew install mvnx
```

### Manual Installation

**For Apple Silicon (M1/M2/M3):**

```bash
# Download
curl -LO https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_darwin_arm64.tar.gz

# Extract
tar -xzf mvnx_darwin_arm64.tar.gz

# Install
sudo mv mvnx /usr/local/bin/

# Clean up
rm mvnx_darwin_arm64.tar.gz
```

**For Intel Macs:**

```bash
# Download
curl -LO https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_darwin_amd64.tar.gz

# Extract
tar -xzf mvnx_darwin_amd64.tar.gz

# Install
sudo mv mvnx /usr/local/bin/

# Clean up
rm mvnx_darwin_amd64.tar.gz
```

---

## Linux

### Quick Install (Recommended)

Use the installation script:

```bash
curl -fsSL https://raw.githubusercontent.com/elitonkfogaca/mvnx-cli/main/install.sh | bash
```

### Manual Installation

**For x86_64:**

```bash
# Download
curl -LO https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_linux_amd64.tar.gz

# Extract
tar -xzf mvnx_linux_amd64.tar.gz

# Install
sudo mv mvnx /usr/local/bin/

# Clean up
rm mvnx_linux_amd64.tar.gz
```

**For ARM64:**

```bash
# Download
curl -LO https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_linux_arm64.tar.gz

# Extract
tar -xzf mvnx_linux_arm64.tar.gz

# Install
sudo mv mvnx /usr/local/bin/

# Clean up
rm mvnx_linux_arm64.tar.gz
```

### Alternative: User-level Installation (No sudo)

If you don't have sudo access:

```bash
# Create local bin directory if it doesn't exist
mkdir -p ~/.local/bin

# Download and extract (example for x86_64)
curl -L https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_linux_amd64.tar.gz | tar xz

# Move to local bin
mv mvnx ~/.local/bin/

# Add to PATH (add this to your ~/.bashrc or ~/.zshrc)
export PATH="$HOME/.local/bin:$PATH"
```

---

## Windows

### PowerShell Installation

Open PowerShell as Administrator and run:

```powershell
# Download the latest release
Invoke-WebRequest -Uri "https://github.com/elitonkfogaca/mvnx-cli/releases/latest/download/mvnx_windows_amd64.zip" -OutFile "mvnx.zip"

# Extract
Expand-Archive -Path mvnx.zip -DestinationPath . -Force

# Move to System32 (requires admin)
Move-Item -Force mvnx.exe C:\Windows\System32\

# Clean up
Remove-Item mvnx.zip
```

### Manual Installation

1. Download the latest Windows release:
   - Go to: https://github.com/elitonkfogaca/mvnx-cli/releases/latest
   - Download: `mvnx_windows_amd64.zip`

2. Extract the ZIP file

3. Move `mvnx.exe` to a directory in your PATH:
   - Recommended: `C:\Windows\System32\` (requires admin)
   - Alternative: Create a folder like `C:\bin` and add it to PATH

### Adding to PATH (Alternative Method)

If you don't want to copy to System32:

1. Create a directory (e.g., `C:\mvnx`)
2. Copy `mvnx.exe` there
3. Add to PATH:
   - Press `Win + X` → System
   - Click "Advanced system settings"
   - Click "Environment Variables"
   - Under "System Variables", find "Path"
   - Click "Edit" → "New"
   - Add: `C:\mvnx`
   - Click "OK" on all dialogs

---

## From Source

### Requirements

- Go 1.21 or later
- Git

### Installation Steps

```bash
# Clone the repository
git clone https://github.com/elitonkfogaca/mvnx-cli.git
cd mvnx-cli

# Build
make build

# Install (Unix-like systems)
sudo mv mvnx /usr/local/bin/

# Or use Go install
go install github.com/elitonkfogaca/mvnx-cli/cmd/mvnx@latest
```

---

## Verification

After installation, verify mvnx is working:

```bash
# Check if mvnx is in PATH
which mvnx  # Unix-like
where mvnx  # Windows

# Check version
mvnx --version

# Test basic functionality
mvnx --help
```

Expected output:
```
mvnx version 1.0.0
commit: abc123...
built: 2026-02-16
```

---

## Troubleshooting

### "mvnx: command not found" (Unix-like)

The binary is not in your PATH. Either:

1. Install to a directory in PATH: `/usr/local/bin`, `/usr/bin`, `~/.local/bin`
2. Add the installation directory to your PATH:
   ```bash
   # Add to ~/.bashrc or ~/.zshrc
   export PATH="/path/to/mvnx/directory:$PATH"
   ```

### "Permission denied" (Unix-like)

Make sure the binary is executable:

```bash
chmod +x /path/to/mvnx
```

### SSL/Certificate errors

If you get certificate errors when downloading:

```bash
# Use --insecure flag (not recommended for production)
curl --insecure -LO <url>

# Or update your CA certificates
# Ubuntu/Debian
sudo apt-get update && sudo apt-get install ca-certificates

# macOS
# Update your system
```

### Windows SmartScreen warning

Windows may show a warning because mvnx is not signed. To proceed:
1. Click "More info"
2. Click "Run anyway"

This is safe for open-source software downloaded from GitHub releases.

### "Access denied" (Windows)

Run PowerShell as Administrator when installing to System32.

Alternatively, install to a user directory and add it to PATH.

---

## Updating mvnx

### Homebrew

```bash
brew upgrade mvnx
```

### Manual

Download and install the latest release following the installation instructions above.

### From Source

```bash
cd mvnx-cli
git pull
make build
sudo make install
```

---

## Uninstalling mvnx

### Homebrew

```bash
brew uninstall mvnx
```

### Manual (Unix-like)

```bash
sudo rm /usr/local/bin/mvnx
# or
rm ~/.local/bin/mvnx
```

### Manual (Windows)

```powershell
Remove-Item C:\Windows\System32\mvnx.exe
# or remove from custom directory
```

---

## Getting Help

- **Issues:** https://github.com/elitonkfogaca/mvnx-cli/issues
- **Documentation:** https://github.com/elitonkfogaca/mvnx-cli/blob/main/README.md
- **Discussions:** https://github.com/elitonkfogaca/mvnx-cli/discussions
