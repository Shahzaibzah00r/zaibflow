package commands

import (
	"context"

	"github.com/Shahzaibzah00r/zaibflow/internal/config"
	"github.com/Shahzaibzah00r/zaibflow/internal/profiles"
	"github.com/Shahzaibzah00r/zaibflow/internal/runtime"
)

func RunLauncher(ctx context.Context, paths config.Paths, secrets config.Secrets, target profiles.Target, args []string, noBanner bool) (int, error) {
	env, err := runtime.BuildEnv(target, secrets)
	if err != nil {
		return 1, err
	}
	return runtime.Launch(ctx, paths, target, args, env, runtime.RunOptions{NoBanner: noBanner})
}
