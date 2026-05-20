//go:build windows

package runtime

import (
	"context"
	"fmt"
	"os/exec"
)

func EnsureClaude(ctx context.Context) error {
	if _, err := exec.LookPath("claude"); err == nil {
		return nil
	}
	return fmt.Errorf("Claude Code CLI is required.\nPlease install it from: https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview")
}
