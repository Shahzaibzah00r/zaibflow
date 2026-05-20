# Changelog

All notable changes to this project will be documented in this file.

## [v0.1.0] - 2026-05-20

### Added

- Full cross-platform support (macOS, Linux, Windows)
- ZaibFlow CLI rebrand from Clother
- `zaibflow run <provider> [args...]` gateway command
- Native Windows support with PowerShell installer (scripts/install.ps1)
- Windows .cmd launcher shims for quick provider access
- npx installer wrapper for easy installation
- GitHub Actions workflow for automated multi-platform builds and releases
- `--bp` flag as alias for `--yolo` (bypass permissions prompt)
- `golang.org/x/term` for cross-platform hidden input (PromptSecret)
- Support for OpenRouter aliases, custom providers, and local models
- Asset logo in assets/logo.png

### Changed

- Renamed all references from "clother" to "ZaibFlow"
- Module path changed to `github.com/Shahzaibzah00r/zaibflow`
- Removed dependency on `/dev/tty` and `stty` command
- Replaced platform-specific TTY handling with golang.org/x/term
- Updated README with comprehensive installation and usage instructions

### Fixed

- PromptSecret now works on Windows using golang.org/x/term.ReadPassword
- CLI argument normalization for permission bypass flag
- GitHub Actions workflow to use modern action-gh-release

### Security

- PromptSecret properly hides sensitive input across all platforms
- Secrets stored with chmod 0o600 file permissions

## License

ZaibFlow is based on the open-source Clother project by jolehuit, licensed under MIT.
