package cli

import (
	"fmt"
	"io"
	"sort"

	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
	"github.com/Shahzaibzah00r/zaibflow/internal/version"
)

func ShowBrief(w io.Writer) {
	fmt.Fprintf(w, "ZaibFlow v%s - Agentic AI Runtime for Claude Code\n\n", version.Value)
	fmt.Fprintln(w, "Usage: zaibflow [options] <command>")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Commands:")
	fmt.Fprintln(w, "  config              Configure a provider")
	fmt.Fprintln(w, "  run <provider>      Run Claude Code through a provider")
	fmt.Fprintln(w, "  list                List profiles")
	fmt.Fprintln(w, "  info                Provider details")
	fmt.Fprintln(w, "  test                Test providers")
	fmt.Fprintln(w, "  status              Show installation state")
	fmt.Fprintln(w, "  update              Update ZaibFlow")
	fmt.Fprintln(w, "  uninstall           Remove ZaibFlow")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Shortcuts:")
	fmt.Fprintln(w, "  zaibflow kimi --bp")
	fmt.Fprintln(w, "  zaibflow zai --bp")
	fmt.Fprintln(w, "  zaibflow openrouter <alias> --bp")
	fmt.Fprintln(w, "  zf-kimi --bp")
	fmt.Fprintln(w, "  zf-or --bp")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Run zaibflow --help for full help.")
}

func ShowFull(w io.Writer, catalog providers.Catalog) {
	fmt.Fprintf(w, "ZaibFlow v%s - Agentic AI Runtime for Claude Code\n", version.Value)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  zaibflow [options] <command> [args]")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Commands:")
	fmt.Fprintln(w, "  config [provider]")
	fmt.Fprintln(w, "  run <provider> [args...]")
	fmt.Fprintln(w, "  list")
	fmt.Fprintln(w, "  info <provider>")
	fmt.Fprintln(w, "  test [provider]")
	fmt.Fprintln(w, "  status")
	fmt.Fprintln(w, "  update")
	fmt.Fprintln(w, "  uninstall")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -h, --help")
	fmt.Fprintln(w, "  -V, --version")
	fmt.Fprintln(w, "  -v, --verbose")
	fmt.Fprintln(w, "  -d, --debug")
	fmt.Fprintln(w, "  -q, --quiet")
	fmt.Fprintln(w, "  -y, --yes")
	fmt.Fprintln(w, "  --bin-dir <path>")
	fmt.Fprintln(w, "  --no-input")
	fmt.Fprintln(w, "  --no-banner")
	fmt.Fprintln(w, "  --json")
	fmt.Fprintln(w, "  --plain")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Launcher shortcuts:")
	fmt.Fprintln(w, "  zf-kimi --bp             skip permission prompts")
	fmt.Fprintln(w, "  zf-zai --bp              skip permission prompts")
	fmt.Fprintln(w, "  zf-or <alias> --bp       skip permission prompts")
	fmt.Fprintln(w, "  zf-local --bp            skip permission prompts")
	fmt.Fprintln(w, "  --yolo, --bp             shorthand for --dangerously-skip-permissions")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Providers:")
	for _, category := range catalog.Categories() {
		fmt.Fprintf(w, "  %s\n", category)
		providersInCategory := catalog.ProvidersByCategory(category)
		sort.SliceStable(providersInCategory, func(i, j int) bool {
			return providersInCategory[i].ID < providersInCategory[j].ID
		})
		for _, provider := range providersInCategory {
			fmt.Fprintf(w, "    %-12s %s\n", provider.ID, provider.Description)
		}
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Advanced:")
	fmt.Fprintln(w, "    openrouter   100+ models via native API")
	fmt.Fprintln(w, "    custom       Anthropic-compatible endpoint")
}
