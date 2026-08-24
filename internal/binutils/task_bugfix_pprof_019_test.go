package binutils

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof019SourceContract(t *testing.T) {
    source, err := os.ReadFile("binutils.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return verMajor > 10 || (verMajor == 10 && verPatch >= 1)") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return verMajor > 10 || (verMajor == 10 || verPatch >= 1)") {
        t.Fatalf("mutated source contract is still present")
    }
}
