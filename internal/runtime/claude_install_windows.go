//go:build windows

package runtime

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Shahzaibzah00r/zaibflow/internal/config"
)

func EnsureClaude(ctx context.Context, paths config.Paths) error {
	if _, err := FindRealClaude(paths); err == nil {
		return nil
	}

	fmt.Fprintln(os.Stderr, "Claude Code CLI is required.")
	fmt.Fprintln(os.Stderr, "Installing Claude Code CLI automatically...")

	// Try PowerShell first
	cmd := exec.CommandContext(ctx, "powershell", "-ExecutionPolicy", "Bypass", "-Command", "irm https://claude.ai/install.ps1 | iex")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		// Fallback to curl + cmd
		cmd = exec.CommandContext(ctx, "cmd", "/c", "curl -fsSL https://claude.ai/install.cmd -o %TEMP%\claude-install.cmd && %TEMP%\claude-install.cmd && del %TEMP%\claude-install.cmd")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	}

	if err != nil {
		return fmt.Errorf("Claude Code CLI installation failed: %w\n\nPlease install manually:\n  PowerShell: irm https://claude.ai/install.ps1 | iex\n  CMD:        curl -fsSL https://claude.ai/install.cmd -o install.cmd && install.cmd && del install.cmd", err)
	}

	if _, err := FindRealClaude(paths); err == nil {
		return nil
	}

	// Claude may have been installed to ~/.local/bin but not on PATH yet.
	home, _ := os.UserHomeDir()
	if home != "" {
		localBin := filepath.Join(home, ".local", "bin")
		localBinClaude := filepath.Join(localBin, "claude.exe")
		if info, err := os.Stat(localBinClaude); err == nil && !info.IsDir() {
			currentPath := os.Getenv("PATH")
			if !strings.Contains(strings.ToLower(currentPath), strings.ToLower(localBin)) {
				os.Setenv("PATH", localBin+string(filepath.ListSeparator)+currentPath)
			}
			if _, err := exec.LookPath("claude"); err == nil {
				return nil
			}
		}
	}

	return fmt.Errorf("Claude Code CLI installed but not found on PATH.\nPlease add %%USERPROFILE%%\\.local\\bin to your PATH and restart your terminal.")
}
