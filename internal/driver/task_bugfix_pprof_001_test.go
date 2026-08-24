package driver

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof001SourceContract(t *testing.T) {
    source, err := os.ReadFile("interactive.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if name == \"sample_index\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
