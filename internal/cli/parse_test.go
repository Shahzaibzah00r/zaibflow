package cli

import (
	"testing"
)

func TestParseAllowsUnknownFlagsAfterCommand(t *testing.T) {
	t.Parallel()

	parsed, err := Parse([]string{"run", "kimi", "--bp"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.Command != "run" {
		t.Fatalf("command = %q, want run", parsed.Command)
	}
	want := []string{"kimi", "--bp"}
	if len(parsed.Args) != len(want) {
		t.Fatalf("args = %v, want %v", parsed.Args, want)
	}
	for i := range want {
		if parsed.Args[i] != want[i] {
			t.Fatalf("args[%d] = %q, want %q", i, parsed.Args[i], want[i])
		}
	}
}

func TestParseAllowsYoloAfterCommand(t *testing.T) {
	t.Parallel()

	parsed, err := Parse([]string{"run", "kimi", "--yolo"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.Command != "run" {
		t.Fatalf("command = %q, want run", parsed.Command)
	}
	want := []string{"kimi", "--yolo"}
	if len(parsed.Args) != len(want) {
		t.Fatalf("args = %v, want %v", parsed.Args, want)
	}
}

func TestParseRejectsUnknownFlagBeforeCommand(t *testing.T) {
	t.Parallel()

	_, err := Parse([]string{"--unknown-flag", "run", "kimi"})
	if err == nil {
		t.Fatal("expected error for unknown flag before command")
	}
}

func TestParseHandlesProviderShortcutWithBp(t *testing.T) {
	t.Parallel()

	parsed, err := Parse([]string{"kimi", "--bp"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.Command != "kimi" {
		t.Fatalf("command = %q, want kimi", parsed.Command)
	}
	if len(parsed.Args) != 1 || parsed.Args[0] != "--bp" {
		t.Fatalf("args = %v, want [--bp]", parsed.Args)
	}
}
