package profile

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisPprof008SourceContract(t *testing.T) {
    source, err := os.ReadFile("merge.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(pm.mappings) == 0 && len(src.Mapping) > 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
