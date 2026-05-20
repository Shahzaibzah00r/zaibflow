# ZaibFlow Installation Guide

Complete installation instructions for all platforms and methods.

## Quick Start

### macOS / Linux (Shell Script)

```bash
# Download and run installer
curl -sSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/zaibflow.sh | bash

# Verify installation
zaibflow --version

# Create launchers (optional - for quick provider access)
zaibflow install
```

### Windows (PowerShell)

```powershell
# Run installer (administrator recommended)
powershell -Command "irm https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex"

# Verify installation
zaibflow --version

# Create launchers (from admin PowerShell)
zaibflow install
```

### npm / npx

```bash
# Global installation
npm install -g zaibflow

# Or use directly without installing
npx zaibflow --help

# Verify
zaibflow --version
```

### Go (from source)

```bash
# Requires Go 1.20+
go install github.com/Shahzaibzah00r/zaibflow/cmd/zaibflow@latest

# Verify
zaibflow --version
```

### Manual Download

1. **Visit** [GitHub Releases](https://github.com/Shahzaibzah00r/zaibflow/releases)
2. **Download** the binary for your OS:
   - `zaibflow_linux_amd64.tar.gz` (Linux)
   - `zaibflow_darwin_amd64.tar.gz` (macOS)
   - `zaibflow_windows_amd64.zip` (Windows)
3. **Extract** the archive
4. **Move** `zaibflow` / `zaibflow.exe` to a directory in your PATH:
   - Linux/macOS: `~/.local/bin/` or `/usr/local/bin/`
   - Windows: `C:\Program Files\ZaibFlow\bin\` (or similar)
5. **Make executable** (Linux/macOS): `chmod +x zaibflow`

## Configuration

### Environment Variables

```bash
# Override config directory (default: ~/.local/share/zaibflow)
export ZAIBFLOW_CONFIG_HOME=~/.config/zaibflow

# Override secrets file location
export ZAIBFLOW_SECRETS_FILE=~/.config/zaibflow/secrets.env

# Disable updates check
export ZAIBFLOW_SKIP_UPDATE_CHECK=1
```

### Configuration File

Create `~/.local/share/zaibflow/config.json`:

```json
{
  "providers": {
    "kimi": {
      "type": "http",
      "baseUrl": "https://api.moonshot.cn/v1",
      "model": "moonshot-v1-8k"
    },
    "openrouter": {
      "type": "http",
      "baseUrl": "https://openrouter.ai/api/v1",
      "apiKeyEnv": "OPENROUTER_API_KEY"
    }
  },
  "aliases": {
    "my-provider": "openrouter"
  }
}
```

## First Run

```bash
# Show available commands
zaibflow --help

# List installed launchers
zaibflow list

# Test with a provider
zaibflow run kimi --version

# Use a launcher shortcut
zf-kimi --version
```

## Creating Launchers

Launchers are quick-access shortcuts to specific providers:

```bash
# Install default launchers
zaibflow install

# List launchers
zaibflow list

# Use a launcher
zf-kimi              # runs: zaibflow run kimi
zf-or               # runs: zaibflow run openrouter
zf-zai              # runs: zaibflow run zai
zf-local            # runs: zaibflow run ollama
zf-custom           # runs: zaibflow run custom
```

## Usage Examples

```bash
# Run with specific provider
zaibflow run kimi --model moonshot-v1-8k

# Run with OpenRouter alias
zaibflow run openrouter my-alias

# Skip permission prompts
zaibflow run kimi --bp
zf-kimi --bp

# Use launcher shortcuts
zf-zai --help
zf-local --version
```

## Troubleshooting

### Permission Denied

**Linux/macOS:**
```bash
# Make binary executable
chmod +x ~/.local/bin/zaibflow
```

**Windows:**
- Right-click PowerShell → "Run as administrator"
- Re-run installer

### Command Not Found

**Add to PATH:**

Linux/macOS:
```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH="$HOME/.local/bin:$PATH"
source ~/.bashrc  # or ~/.zshrc
```

**Windows:**
- Search "Environment Variables"
- Add `C:\Program Files\ZaibFlow\bin` to System PATH
- Restart PowerShell/Terminal

### Provider Connection Issues

```bash
# Check configuration
zaibflow info

# Verify provider
zaibflow run <provider> --help

# Test connection
zaibflow run <provider> --version
```

### Clear Cache / Reset

```bash
# Remove launchers
zaibflow uninstall

# Clear config (backup first!)
rm -rf ~/.local/share/zaibflow

# Reinstall
zaibflow install
```

## Platform-Specific Notes

### macOS

- Requires macOS 10.12+
- Use Homebrew for easy updates: `brew install Shahzaibzah00r/tap/zaibflow` (when published)
- Terminal must support dark mode for hidden input prompts

### Linux

- Tested on Ubuntu 20.04+, Debian 10+, Fedora 30+
- May require `libterm` for terminal control (usually pre-installed)
- Use system package manager if available

### Windows

- Requires Windows 10+, PowerShell 5.0+
- Windows Defender may require adding exclusion for extraction folder
- Administrator privileges recommended for installation

## Support

**Issues:** [GitHub Issues](https://github.com/Shahzaibzah00r/zaibflow/issues)
**Email:** shahzaibzahoor7@gmail.com

## License

ZaibFlow is based on the open-source Clother project by jolehuit, licensed under MIT.
