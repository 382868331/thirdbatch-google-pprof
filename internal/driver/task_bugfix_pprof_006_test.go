package driver

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof006SourceContract(t *testing.T) {
    source, err := os.ReadFile("driver.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer cleanupTempFiles()") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "cleanupTempFiles()") {
        t.Fatalf("mutated source contract is still present")
    }
}
