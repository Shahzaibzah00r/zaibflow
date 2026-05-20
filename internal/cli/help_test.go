package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Shahzaibzah00r/zaibflow/internal/providers"
)

func TestShowBriefDoesNotContainClother(t *testing.T) {
	var buf bytes.Buffer
	ShowBrief(&buf)
	out := buf.String()
	if strings.Contains(out, "Clother") {
		t.Fatalf("ShowBrief contains old branding 'Clother':\n%s", out)
	}
	if !strings.Contains(out, "ZaibFlow") {
		t.Fatal("ShowBrief missing ZaibFlow branding")
	}
}

func TestShowFullDoesNotContainClother(t *testing.T) {
	catalog, err := providers.Load()
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	ShowFull(&buf, catalog)
	out := buf.String()
	if strings.Contains(out, "Clother") {
		t.Fatalf("ShowFull contains old branding 'Clother':\n%s", out)
	}
	if !strings.Contains(out, "ZaibFlow") {
		t.Fatal("ShowFull missing ZaibFlow branding")
	}
}
