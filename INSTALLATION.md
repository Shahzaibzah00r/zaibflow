# ZaibFlow Installation Guide

## Quick Start

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.sh | bash
```

The installer will:
- Detect your OS and architecture
- Download the latest ZaibFlow release
- Install to `$HOME/bin` (or `$HOME/.local/bin` as fallback)
- Add the install directory to your PATH automatically
- Create launcher shortcuts (`zf-kimi`, `zf-zai`, `zf-or`, `zf-local`)
- Check for Claude Code CLI and guide you if it's missing

### Windows

```powershell
irm https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex
```

The installer will:
- Detect your architecture
- Download the latest ZaibFlow release
- Install to `%LOCALAPPDATA%\Programs\ZaibFlow\bin`
- Add the install directory to your User PATH automatically
- Create `.cmd` launcher shortcuts
- Check for Claude Code CLI and guide you if it's missing

### npm / npx

```bash
npx zaibflow
```

If ZaibFlow is not already installed, the npm wrapper will download and install it automatically, then forward your arguments.

### Go (from source)

```bash
go install github.com/Shahzaibzah00r/zaibflow/cmd/zaibflow@latest
```

### Manual Download

1. Visit [GitHub Releases](https://github.com/Shahzaibzah00r/zaibflow/releases)
2. Download the archive for your OS:
   - `zaibflow_linux_amd64.tar.gz`
   - `zaibflow_linux_arm64.tar.gz`
   - `zaibflow_darwin_amd64.tar.gz`
   - `zaibflow_darwin_arm64.tar.gz`
   - `zaibflow_windows_amd64.zip`
   - `zaibflow_windows_arm64.zip`
3. Extract the archive — it contains a single `zaibflow` (or `zaibflow.exe`) binary
4. Move it to a directory on your PATH

## First Run

```bash
# Configure a provider
zaibflow config

# Run with permission bypass
zaibflow kimi --bp
zaibflow zai --bp
zaibflow openrouter <alias> --bp

# Use launcher shortcuts
zf-kimi --bp
zf-zai --bp
zf-or <alias> --bp
zf-local --bp
```

## Environment Variables

```bash
# Override config directory
export ZAIBFLOW_CONFIG_DIR=~/.config/zaibflow

# Override bin directory
export ZAIBFLOW_BIN=~/bin

# Disable update check
export ZAIBFLOW_SKIP_UPDATE_CHECK=1
```

## Troubleshooting

### Command Not Found

Restart your terminal so PATH changes take effect.

### Provider Connection Issues

```bash
# Check provider info
zaibflow info <provider>

# Test provider
zaibflow test <provider>
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

## Support

**Issues:** [GitHub Issues](https://github.com/Shahzaibzah00r/zaibflow/issues)
**Email:** <shahzaibzahoor7@gmail.com>

## License

ZaibFlow is based on the open-source Clother project by jolehuit, licensed under MIT.
