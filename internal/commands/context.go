package commands

import (
	"github.com/Shahzaibzah00r/zaibflow/internal/cli"
	"github.com/Shahzaibzah00r/zaibflow/internal/config"
	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
	"github.com/Shahzaibzah00r/zaibflow/internal/ui"
)

type Context struct {
	Paths   config.Paths
	Config  *config.File
	Secrets config.Secrets
	Catalog providers.Catalog
	Output  *ui.Output
	Prompt  *ui.Prompter
	Options cli.Options
}
