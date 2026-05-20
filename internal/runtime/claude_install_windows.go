//go:build windows

package runtime

import (
	"context"
	"fmt"

	"github.com/Shahzaibzah00r/zaibflow/internal/config"
)

func EnsureClaude(ctx context.Context, paths config.Paths) error {
	if _, err := FindRealClaude(paths); err == nil {
		return nil
	}
	return fmt.Errorf("Claude Code CLI is required.\nPlease install it from: https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview")
}
