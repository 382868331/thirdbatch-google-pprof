package driver

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof003SourceContract(t *testing.T) {
    source, err := os.ReadFile("options.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if d.Sym == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
