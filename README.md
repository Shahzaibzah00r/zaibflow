<div align="center">
  <img src="assets/logo.png" alt="ZaibFlow logo" width="220" />
  <h1>ZaibFlow</h1>
  <p><strong>A cross-platform CLI to manage and invoke model providers (Claude/OpenRouter/OLLAMA/custom).</strong></p>
  <p>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT License" /></a>
    <a href="https://go.dev/"><img src="https://img.shields.io/badge/Language-Go-00ADD8.svg" alt="Go" /></a>
    <a href="#platform-support"><img src="https://img.shields.io/badge/Platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey.svg" alt="Platform support" /></a>
  </p>
</div>

## Quick Install

Choose one of the following installation methods based on your platform and preferences.

- Install via the official bootstrap script (recommended):

```bash
curl -fsSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/zaibflow.sh | bash
```

- Windows (PowerShell installer):

```powershell
iwr https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex
```

- Install with `go` (build from source):

```bash
go install github.com/Shahzaibzah00r/zaibflow@latest
```

- Run directly with `npx` (uses packaged wrapper):

```bash
npx zaibflow
```

- Manual download for Windows: grab the latest `zaibflow_windows_amd64.zip` from the GitHub Releases page, extract and place `zaibflow.exe` on your PATH.

## Platform-specific notes

- macOS
  - Recommended: run the bootstrap script above. If you use Homebrew and a tap is available you can `brew install zaibflow` (tap not provided here by default).
  - After installation, ensure `~/.local/bin` or the installed `bin` directory is on your `PATH`.

- Linux
  - Use the bootstrap script, or install via `go install` if you have Go toolchain available.
  - Place the `zaibflow` binary in a directory on your `$PATH` (for single-user installs, `~/.local/bin` is a common choice).

- Windows
  - Use the PowerShell installer above to download the packaged `.zip` and create the shim.
  - Alternatively, download the binary release, extract `zaibflow.exe`, and add its folder to your `%PATH%`.
  - The repository includes `scripts/install.ps1` for unattended installs.

## Building from source

Prerequisites: Go 1.20+ installed.

```bash
git clone https://github.com/Shahzaibzah00r/zaibflow.git
cd zaibflow
go build ./cmd/zaibflow
# or
go install github.com/Shahzaibzah00r/zaibflow/cmd/zaibflow@latest
```

## Usage examples

Basic commands:

```bash
zaibflow init            # create config & setup default files
zaibflow config edit     # edit configuration
zaibflow run <provider> [args...]   # run a provider-based launcher
zaibflow install         # install launchers and helpers
zaibflow update          # check for updates and install
```

Examples:

```bash
zaibflow run ollama --model qwen3-coder
zaibflow run openrouter my-alias
zaibflow run local-provider --help
```

Launchers

ZaibFlow creates small shims/launchers to make invoking providers easy. Typical names:

- `zf-<provider>` (recommended short form)
- `zaibflow-<provider>` (legacy compatibility)

These are installed into the configured `bin` directory (see `ZAIBFLOW_BIN`).

## Environment variables

- `ZAIBFLOW_BIN` — custom bin directory where launchers are written.
- `ZAIBFLOW_CONFIG_DIR` — config directory (defaults to XDG config path + `/zaibflow`).
- `ZAIBFLOW_DATA_DIR` — data directory (defaults to XDG data path + `/zaibflow`).
- `ZAIBFLOW_CACHE_DIR` — cache directory.
- `ZAIBFLOW_SKIP_SELF_UPDATE=1` — disable automatic self-update during `install`.
- `ZAIBFLOW_RELEASE_BASE_URL` — override release base URL to an alternate host.

## Troubleshooting

- If `zaibflow` is not found after install, ensure the install `bin` directory is on your `PATH`.
- If provider launchers are missing, run `zaibflow install` to re-generate shims.

## Contributing

Contributions welcome: open issues and PRs against the GitHub repository.

## Contact

Maintained by Shahzaib — contact: <shahzaibzahoor7@gmail.com>

## License

MIT
 
ZaibFlow is based on the open-source Clother project by jolehuit, licensed under MIT.
