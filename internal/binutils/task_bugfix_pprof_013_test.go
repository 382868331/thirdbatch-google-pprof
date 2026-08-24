package binutils

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof013SourceContract(t *testing.T) {
    source, err := os.ReadFile("disasm.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if symAddr == start {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if symAddr != start {") {
        t.Fatalf("mutated source contract is still present")
    }
}
