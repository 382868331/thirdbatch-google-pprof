package binutils

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof005SourceContract(t *testing.T) {
    source, err := os.ReadFile("binutils.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer f.Close()") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "f.Close()") {
        t.Fatalf("mutated source contract is still present")
    }
}
