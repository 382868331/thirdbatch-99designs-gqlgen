package fieldset

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen003SourceContract(t *testing.T) {
    source, err := os.ReadFile("fieldset.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if def == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
