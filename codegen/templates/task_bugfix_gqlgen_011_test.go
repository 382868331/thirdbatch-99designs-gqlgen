package templates

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen011SourceContract(t *testing.T) {
    source, err := os.ReadFile("import.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "alias = aliases[0]") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "alias = aliases[1]") {
        t.Fatalf("mutated source contract is still present")
    }
}
