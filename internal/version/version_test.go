package version

import (
	"strings"
	"testing"
)

func TestVersionIsNotOldClotherVersion(t *testing.T) {
	if strings.Contains(Value, "3.0.9") {
		t.Fatalf("version.Value = %q, should not contain old Clother version 3.0.9", Value)
	}
}

func TestVersionIsZeroOneZero(t *testing.T) {
	if Value != "0.1.0" {
		t.Fatalf("version.Value = %q, want 0.1.0", Value)
	}
}
