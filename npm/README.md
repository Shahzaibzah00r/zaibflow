# ZaibFlow

ZaibFlow is an **Agentic AI Runtime for Claude Code** — a cross-platform CLI that lets you run Claude Code with any AI provider. Switch between Kimi, Z.AI, OpenRouter, Ollama, and 10+ providers with a single command.

## Quick Start

```bash
npx zaibflow
```

Or install globally so `zaibflow` and `zf` are always available:

```bash
npm install -g zaibflow
```

## Usage

Configure your favorite provider:

```bash
zaibflow config
```

Run Claude Code through a provider:

```bash
zaibflow kimi --bp           # Kimi K2
zaibflow zai --bp            # Z.AI
zaibflow openrouter <alias> --bp
zaibflow ollama --bp
```

Use launcher shortcuts:

```bash
zf --help                    # Same as zaibflow
zf-kimi --bp                 # Quick launch Kimi
zf-zai --bp
zf-or <alias> --bp
zf-local --bp
```

## What is ZaibFlow?

ZaibFlow is a CLI tool and launcher ecosystem for Claude Code users who want flexibility across AI providers. Instead of being locked into a single model, you can:

- Run Claude Code with **Kimi K2**, **Z.AI**, **OpenRouter**, **Ollama**, **DeepSeek**, **MiniMax**, and more
- Skip permission prompts with `--bp` or `--yolo`
- Use short launcher commands like `zf-kimi` instead of typing full paths
- Auto-install Claude Code CLI if it's missing

## Supported Providers

- **Kimi** (kimi-k2.5)
- **Z.AI** (glm-5)
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

- [Claude Code CLI](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview) — ZaibFlow checks for it automatically and installs it if missing (macOS/Linux)
- Node.js 18+

## Platform Support

- macOS
- Linux
- Windows

## Install Methods

- **macOS/Linux**: `curl -fsSL https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.sh | bash`
- **Windows**: `irm https://raw.githubusercontent.com/Shahzaibzah00r/zaibflow/main/scripts/install.ps1 | iex`
- **npm**: `npm install -g zaibflow`
- **npx**: `npx zaibflow`

## Links

- [GitHub](https://github.com/Shahzaibzah00r/zaibflow)
- [Issues](https://github.com/Shahzaibzah00r/zaibflow/issues)

## License

MIT
