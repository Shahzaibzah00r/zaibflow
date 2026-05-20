package runtime

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"strings"
	"time"

	"github.com/Shahzaibzah00r/zaibflow/internal/config"
	"github.com/Shahzaibzah00r/zaibflow/internal/session"
	"github.com/Shahzaibzah00r/zaibflow/internal/update"
	"github.com/Shahzaibzah00r/zaibflow/internal/version"
)

func RunClaudeShim(ctx context.Context, paths config.Paths, args []string) (int, error) {
	args = NormalizeClaudeArgs(args)
	if isTTY(os.Stderr) && !IsHomebrew() {
		if message, err := update.MaybeMessage(paths, version.Value, time.Now()); err == nil && message != "" {
			fmt.Fprintln(os.Stderr, message)
		}
	}
	claudePath, err := FindRealClaude(paths)
	if err != nil {
		return 1, err
	}
	if err := session.RestoreStale(paths); err != nil {
		return 1, err
	}
	if code, handled, err := runWithTemporaryPatch(ctx, claudePath, paths, args, os.Environ(), ""); handled {
		return code, err
	}
	return runClaudeCommand(ctx, claudePath, args, os.Environ(), "")
}

func FindRealClaude(paths config.Paths) (string, error) {
	self, _ := os.Executable()
	selfResolved := resolvedPath(self)
	if goruntime.GOOS == "windows" {
		candidate, err := exec.LookPath("claude")
		if err == nil && candidate != "" {
			if selfResolved == "" || !samePath(candidate, selfResolved) {
				return candidate, nil
			}
		}
		fallback := filepath.Join(paths.BinDir, "claude-real")
		if info, err := os.Stat(fallback); err == nil && !info.IsDir() {
			return fallback, nil
		}
		return "", fmt.Errorf("could not locate real claude; ensure `claude` is in PATH or `%s` exists", fallback)
	}
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if dir == "" {
			continue
		}
		candidate := filepath.Join(dir, "claude")
		info, err := os.Stat(candidate)
		if err != nil || info.IsDir() {
			continue
		}
		if selfResolved != "" && samePath(candidate, selfResolved) {
			continue
		}
		return candidate, nil
	}
	fallback := filepath.Join(paths.BinDir, "claude-real")
	if info, err := os.Stat(fallback); err == nil && !info.IsDir() {
		if selfResolved == "" || !samePath(fallback, selfResolved) {
			return fallback, nil
		}
	}
	return "", fmt.Errorf("could not locate real claude; ensure `claude` is in PATH or `%s` exists", fallback)
}

func PreserveRealClaude(paths config.Paths, realClaudePath string) error {
	if realClaudePath == "" {
		return nil
	}
	defaultClaude := filepath.Join(paths.BinDir, "claude")
	if !samePath(realClaudePath, defaultClaude) {
		return nil
	}

	preserved := filepath.Join(paths.BinDir, "claude-real")
	if samePath(defaultClaude, preserved) {
		return nil
	}

	if _, err := os.Stat(preserved); err == nil {
		if err := os.Remove(preserved); err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	return os.Rename(defaultClaude, preserved)
}

func resolvedPath(path string) string {
	if path == "" {
		return ""
	}
	resolved, err := filepath.EvalSymlinks(path)
	if err == nil {
		path = resolved
	}
	abs, err := filepath.Abs(path)
	if err == nil {
		path = abs
	}
	return filepath.Clean(path)
}

func samePath(left, right string) bool {
	if left == "" || right == "" {
		return false
	}
	return strings.EqualFold(resolvedPath(left), resolvedPath(right))
}
