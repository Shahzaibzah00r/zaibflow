package commands

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
	"github.com/Shahzaibzah00r/zaibflow/internal/ui"
)

func TestDispatchUnknownCommandShowsHelpfulError(t *testing.T) {
	t.Parallel()

	catalog, err := providers.Load()
	if err != nil {
		t.Fatal(err)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	ctx := Context{
		Catalog: catalog,
		Output:  &ui.Output{Stdout: &stdout, Stderr: &stderr, Format: ui.FormatHuman},
	}

	_, err = Dispatch(context.Background(), ctx, "unknownprovider", []string{})
	if err == nil {
		t.Fatal("expected error for unknown provider")
	}
	msg := err.Error()
	if !strings.Contains(msg, "unknown command or provider") {
		t.Fatalf("error message does not contain 'unknown command or provider': %q", msg)
	}
	if !strings.Contains(msg, "zaibflow config") {
		t.Fatalf("error message does not suggest 'zaibflow config': %q", msg)
	}
}
