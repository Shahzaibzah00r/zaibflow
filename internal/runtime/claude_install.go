//go:build !windows

package runtime

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func EnsureClaude(ctx context.Context) error {
	if _, err := exec.LookPath("claude"); err == nil {
		return nil
	}

	fmt.Fprintln(os.Stderr, "Claude Code CLI is required.")
	fmt.Fprintln(os.Stderr, "Installing Claude Code CLI automatically...")
	cmd := exec.CommandContext(ctx, "bash", "-c", "curl -fsSL https://claude.ai/install.sh | bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Claude Code CLI installation failed: %w", err)
	}
	if _, err := exec.LookPath("claude"); err != nil {
		return fmt.Errorf("Claude Code CLI installed but not found on PATH; please restart your terminal")
	}
	return nil
}
