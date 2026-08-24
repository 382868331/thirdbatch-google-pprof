package profile

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof014SourceContract(t *testing.T) {
    source, err := os.ReadFile("filter.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if focus == nil && ignore == nil && hide == nil && show == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if focus != nil && ignore == nil && hide == nil && show == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
