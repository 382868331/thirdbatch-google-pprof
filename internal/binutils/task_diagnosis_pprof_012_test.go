package binutils

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisPprof012SourceContract(t *testing.T) {
    source, err := os.ReadFile("addr2liner.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(stack) > 0 && d.nm != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
