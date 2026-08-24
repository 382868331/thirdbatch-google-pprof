package binutils

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof002SourceContract(t *testing.T) {
    source, err := os.ReadFile("addr2liner_nm.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if line == \"\" && err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
