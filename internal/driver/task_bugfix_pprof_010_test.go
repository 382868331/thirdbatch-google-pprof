package driver

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof010SourceContract(t *testing.T) {
    source, err := os.ReadFile("driver_focus.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if value == \"\" || err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
