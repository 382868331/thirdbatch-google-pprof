package report

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof009SourceContract(t *testing.T) {
    source, err := os.ReadFile("report.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if na, nb := a.sym.Name[0], b.sym.Name[0]; na != nb {") {
        t.Fatalf("expected source contract is missing")
    }
}
