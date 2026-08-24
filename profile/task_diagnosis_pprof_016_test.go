package profile

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisPprof016SourceContract(t *testing.T) {
    source, err := os.ReadFile("legacy_profile.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return true") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return false") {
        t.Fatalf("mutated source contract is still present")
    }
}
