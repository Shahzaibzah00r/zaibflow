package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/jolehuit/clother/internal/profiles"
)

func runRun(ctx context.Context, c Context, args []string) (int, error) {
	if len(args) == 0 {
		return 1, fmt.Errorf("usage: zaibflow run <provider> [args...]")
	}

	profile := args[0]
	forwarded := args[1:]
	if profile == "openrouter" || profile == "or" {
		if len(forwarded) == 0 || strings.HasPrefix(forwarded[0], "-") {
			return 1, fmt.Errorf("usage: zaibflow run openrouter <alias> [args...]")
		}
		profile = "or-" + forwarded[0]
		forwarded = forwarded[1:]
	} else if profile == "custom" {
		if len(forwarded) == 0 || strings.HasPrefix(forwarded[0], "-") {
			return 1, fmt.Errorf("usage: zaibflow run custom <provider-name> [args...]")
		}
		profile = forwarded[0]
		forwarded = forwarded[1:]
	}

	target, err := profiles.Resolve(profile, c.Catalog, c.Config)
	if err != nil {
		return 1, err
	}
	return RunLauncher(ctx, c.Paths, c.Secrets, target, forwarded, c.Options.NoBanner)
}
