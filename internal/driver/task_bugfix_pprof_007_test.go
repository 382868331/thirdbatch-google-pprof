package driver

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof007SourceContract(t *testing.T) {
    source, err := os.ReadFile("interactive.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if c == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if c != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
