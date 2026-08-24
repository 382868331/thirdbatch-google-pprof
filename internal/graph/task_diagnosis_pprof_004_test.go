package graph

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisPprof004SourceContract(t *testing.T) {
    source, err := os.ReadFile("dotgraph.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "nodeIDMap[n] = i + 1") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "nodeIDMap[n] = i - 1") {
        t.Fatalf("mutated source contract is still present")
    }
}
