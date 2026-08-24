package profile

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixPprof018SourceContract(t *testing.T) {
    source, err := os.ReadFile("merge.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "lines[i*2+1] = strconv.FormatInt(line.Line, 16)") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "lines[i*2- 1] = strconv.FormatInt(line.Line, 16)") {
        t.Fatalf("mutated source contract is still present")
    }
}
