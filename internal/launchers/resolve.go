package launchers

import (
	"fmt"

	"github.com/Shahzaibzah00r/zaibflow/internal/config"
	"github.com/Shahzaibzah00r/zaibflow/internal/profiles"
	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
)

func Resolve(argv0 string, catalog providers.Catalog, cfg *config.File) (profiles.Target, bool, error) {
	profile, ok := profiles.Invocation(argv0)
	if !ok {
		return profiles.Target{}, false, nil
	}
	target, err := profiles.Resolve(profile, catalog, cfg)
	if err != nil {
		return profiles.Target{}, true, fmt.Errorf("resolve launcher %s: %w", profile, err)
	}
	return target, true, nil
}
