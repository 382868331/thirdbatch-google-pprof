package driver

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisPprof020SourceContract(t *testing.T) {
    source, err := os.ReadFile("webui.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if rpt == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
