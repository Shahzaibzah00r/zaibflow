package app

import (
	"testing"

	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
)

func TestIsProviderShortcutKnownProviders(t *testing.T) {
	t.Parallel()

	known := []string{"kimi", "zai", "zai-cn", "minimax", "minimax-cn", "moonshot", "ve", "deepseek", "mimo", "alibaba", "alibaba-us", "alibaba-cn", "ollama", "lmstudio", "llamacpp", "native", "openrouter", "custom"}
	for _, name := range known {
		if !isProviderShortcut(name, providers.Catalog{}) {
			t.Fatalf("isProviderShortcut(%q) = false, want true", name)
		}
	}
}

func TestIsProviderShortcutUnknown(t *testing.T) {
	t.Parallel()

	if isProviderShortcut("", providers.Catalog{}) {
		t.Fatal("isProviderShortcut(\"\") = true, want false")
	}
	if isProviderShortcut("foobar", providers.Catalog{}) {
		t.Fatal("isProviderShortcut(\"foobar\") = true, want false")
	}
}

func TestIsProviderShortcutUsesCatalog(t *testing.T) {
	t.Parallel()

	catalog, err := providers.Load()
	if err != nil {
		t.Fatal(err)
	}

	for _, id := range catalog.IDs() {
		if !isProviderShortcut(id, catalog) {
			t.Fatalf("isProviderShortcut(%q, catalog) = false, want true", id)
		}
	}
}
