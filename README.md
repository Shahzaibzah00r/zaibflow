<div align="center">
  <img src="assets/logo.png" alt="ZaibFlow logo" width="220" />
  <h1>ZaibFlow</h1>
  <p><strong>Agentic AI Runtime for Claude Code — run any AI provider through Claude Code.</strong></p>
  <p>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT License" /></a>
    <a href="https://go.dev/"><img src="https://img.shields.io/badge/Language-Go-00ADD8.svg" alt="Go" /></a>
    <a href="#platform-support"><img src="https://img.shields.io/badge/Platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey.svg" alt="Platform support" /></a>
  </p>
</div>

## One-Command Install

### macOS / Linux

Open your terminal and run:

```bash
curl -fsSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.sh | bash
```

### Windows (PowerShell)

Open **PowerShell** and run:

```powershell
irm https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex
```

> `irm` = `Invoke-RestMethod` (downloads the script)  
> `iex` = `Invoke-Expression` (runs the script)

### npm / npx

Works on any platform with Node.js 18+:

```bash
npx zaibflow
```

Or install globally so `zaibflow` and `zf` are always available:

```bash
npm install -g zaibflow
```

> ZaibFlow requires Claude Code CLI. The installer checks for it automatically and installs it if it's missing (macOS/Linux). On Windows it will guide you if it's missing.

## What is ZaibFlow?

ZaibFlow is a CLI tool and launcher ecosystem for Claude Code users who want flexibility across AI providers. Instead of being locked into a single model, you can:

- Run Claude Code with **Kimi K2**, **Z.AI**, **OpenRouter**, **Ollama**, **DeepSeek**, **MiniMax**, and more
- Skip permission prompts with `--bp` or `--yolo`
- Use short launcher commands like `zf-kimi` instead of typing full paths
- Auto-install Claude Code CLI if it's missing

## Usage

```bash
zaibflow config              # Configure a provider
zaibflow kimi --bp           # Run Kimi K2 with permission bypass
zaibflow zai --bp            # Run Z.A.I with permission bypass
zaibflow openrouter <alias> --bp
zaibflow ollama --bp
zaibflow custom <name> --bp
zf --help                    # Main shortcut, same as zaibflow
zf-kimi --bp                 # Launcher shortcut
zf-zai --bp
zf-or <alias> --bp
zf-local --bp
zf-custom --bp
```

## Supported Providers

- **Kimi** (kimi-k2.5)
- **Z.A.I** (glm-5)
- **OpenRouter** (100+ models)
- **Ollama** (local models)
- **DeepSeek**
- **MiniMax**
- **Moonshot**
- **Alibaba**
- **Xiaomi MiMo**
- **VolcEngine**
- And more...

## Platform-Specific Notes

- **macOS / Linux**: The installer prefers `$HOME/bin`, falling back to `$HOME/.local/bin`. It adds the directory to your shell PATH automatically.
- **Windows**: The installer places binaries in `%LOCALAPPDATA%\Programs\ZaibFlow\bin` and updates your User PATH automatically.
- **npm/npx**: The first run downloads the correct native binary for your platform automatically. No manual PATH setup needed.

## Building from Source

Prerequisites: Go 1.20+

```bash
git clone https://github.com/Shahzaibzah00r/zaibflow.git
cd zaibflow
go build ./cmd/zaibflow
# or
go install github.com/Shahzaibzah00r/zaibflow/cmd/zaibflow@latest
```

## Environment Variables

- `ZAIBFLOW_BIN` — custom bin directory where launchers are written.
- `ZAIBFLOW_CONFIG_DIR` — config directory (defaults to XDG config path + `/zaibflow`).
- `ZAIBFLOW_DATA_DIR` — data directory (defaults to XDG data path + `/zaibflow`).
- `ZAIBFLOW_CACHE_DIR` — cache directory.
- `ZAIBFLOW_SKIP_SELF_UPDATE=1` — disable automatic self-update during `install`.
- `ZAIBFLOW_RELEASE_BASE_URL` — override release base URL to an alternate host.

## Troubleshooting

- If `zaibflow` is not found after install, **restart your terminal** so PATH changes take effect.
- If provider launchers are missing, run `zaibflow install` to re-generate shims.
- If `npx zaibflow` shows a 404, make sure you are using npm 18+ and have a working internet connection.

## Contributing

Contributions welcome: open issues and PRs against the GitHub repository.

## Contact

Maintained by Shahzaib — contact: <shahzaibzahoor7@gmail.com>

## License

MIT

ZaibFlow is based on the open-source Clother project by jolehuit, licensed under MIT.
