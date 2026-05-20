<p align="center">
  <img src="https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/assets/logo.png" width="220" alt="ZaibFlow logo">
</p>

<h1 align="center">ZaibFlow</h1>

<p align="center"><strong>Agentic AI Runtime for Claude Code</strong> — a cross-platform CLI that lets you run Claude Code with any AI provider. Switch between Kimi, Z.A.I, OpenRouter, Ollama, and 10+ providers with a single command.</p>

## Quick Start

```bash
npx zaibflow
```

Or install globally so `zaibflow` and `zf` are always available:

```bash
npm install -g zaibflow
```

## One-Command Install (No npm)

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex
```

> `irm` = `Invoke-RestMethod` (downloads the script)  
> `iex` = `Invoke-Expression` (runs the script)

## Usage

Configure your favorite provider:

```bash
zaibflow config
```

Run Claude Code through a provider:

```bash
zaibflow kimi --bp           # Kimi K2
zaibflow zai --bp            # Z.A.I
zaibflow openrouter <alias> --bp
zaibflow ollama --bp
zaibflow custom <name> --bp
```

Use launcher shortcuts:

```bash
zf --help                    # Same as zaibflow
zf-kimi --bp                 # Quick launch Kimi
zf-zai --bp
zf-or <alias> --bp
zf-local --bp
zf-custom --bp
```

## What is ZaibFlow?

ZaibFlow is a CLI tool and launcher ecosystem for Claude Code users who want flexibility across AI providers. Instead of being locked into a single model, you can:

- Run Claude Code with **Kimi K2**, **Z.A.I**, **OpenRouter**, **Ollama**, **DeepSeek**, **MiniMax**, and more
- Skip permission prompts with `--bp` or `--yolo`
- Use short launcher commands like `zf-kimi` instead of typing full paths
- Auto-install Claude Code CLI if it's missing (macOS/Linux/Windows)

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

## Requirements

- **Claude Code CLI** — ZaibFlow auto-installs it for you on macOS/Linux/Windows. If you prefer to install manually:
  - **macOS/Linux**: `curl -fsSL https://claude.ai/install.sh | bash`
  - **Windows (PowerShell)**: `irm https://claude.ai/install.ps1 | iex`
  - **Windows (CMD)**: `curl -fsSL https://claude.ai/install.cmd -o install.cmd && install.cmd && del install.cmd`
- **Node.js 18+** (only for npm/npx install)

## Platform Support

- macOS
- Linux
- Windows

## Install Methods

- **macOS/Linux**: `curl -fsSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.sh | bash`
- **Windows (PowerShell)**: `irm https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex`
- **npm**: `npm install -g zaibflow`
- **npx**: `npx zaibflow`

## Troubleshooting

- If `zaibflow` is not found after install, **restart your terminal** so PATH changes take effect.
- If you see a 404 error on first run, check your internet connection and try again.
- Run `zaibflow install` to re-generate launcher shims if they are missing.
- On Windows, if Claude Code CLI installs but zaibflow still says it's missing, add `%USERPROFILE%\.local\bin` to your User PATH and restart your terminal.

## Links

- [GitHub](https://github.com/Shahzaibzah00r/zaibflow)
- [Issues](https://github.com/Shahzaibzah00r/zaibflow/issues)

## License

MIT
