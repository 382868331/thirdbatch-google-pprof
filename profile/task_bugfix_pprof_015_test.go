package profile

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof015SourceContract(t *testing.T) {
    source, err := os.ReadFile("merge.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if id < uint64(len(lm.dense)) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if id <= uint64(len(lm.dense)) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
