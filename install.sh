#!/usr/bin/env bash
#
# Install script for mvnx
# Usage: curl -fsSL https://raw.githubusercontent.com/elitonkfogaca/mvnx-cli/main/install.sh | bash
#

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect OS and architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux*)
        OS_TYPE="linux"
        ;;
    Darwin*)
        OS_TYPE="darwin"
        ;;
    *)
        echo -e "${RED}Error: Unsupported operating system: $OS${NC}"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64|amd64)
        ARCH_TYPE="amd64"
        ;;
    arm64|aarch64)
        ARCH_TYPE="arm64"
        ;;
    *)
        echo -e "${RED}Error: Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

# Configuration
REPO="elitonkfogaca/mvnx-cli"
INSTALL_DIR="${MVNX_INSTALL_DIR:-/usr/local/bin}"
BINARY_NAME="mvnx"

echo -e "${GREEN}Installing mvnx...${NC}"
echo "OS: $OS_TYPE"
echo "Architecture: $ARCH_TYPE"
echo "Install directory: $INSTALL_DIR"
echo ""

# Get latest release version
echo -e "${YELLOW}Fetching latest release...${NC}"
LATEST_VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}Error: Could not fetch latest version${NC}"
    exit 1
fi

echo "Latest version: $LATEST_VERSION"

# Download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/${LATEST_VERSION}/mvnx_${OS_TYPE}_${ARCH_TYPE}.tar.gz"

echo -e "${YELLOW}Downloading from: $DOWNLOAD_URL${NC}"

# Create temporary directory
TMP_DIR=$(mktemp -d)
trap 'rm -rf "$TMP_DIR"' EXIT

# Download and extract
cd "$TMP_DIR"
if ! curl -fsSL "$DOWNLOAD_URL" | tar -xz; then
    echo -e "${RED}Error: Failed to download or extract mvnx${NC}"
    exit 1
fi

# Check if binary exists
if [ ! -f "$BINARY_NAME" ]; then
    echo -e "${RED}Error: Binary not found in archive${NC}"
    exit 1
fi

# Make binary executable
chmod +x "$BINARY_NAME"

# Check if install directory needs sudo
if [ -w "$INSTALL_DIR" ]; then
    mv "$BINARY_NAME" "$INSTALL_DIR/"
else
    echo -e "${YELLOW}Installing to $INSTALL_DIR requires sudo...${NC}"
    sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
fi

# Verify installation
if command -v mvnx >/dev/null 2>&1; then
    VERSION=$(mvnx --version 2>/dev/null || echo "unknown")
    echo ""
    echo -e "${GREEN}✓ mvnx installed successfully!${NC}"
    echo ""
    echo "Run 'mvnx --help' to get started"
else
    echo ""
    echo -e "${YELLOW}⚠ mvnx was installed, but is not in your PATH${NC}"
    echo "Add $INSTALL_DIR to your PATH to use mvnx"
fi
