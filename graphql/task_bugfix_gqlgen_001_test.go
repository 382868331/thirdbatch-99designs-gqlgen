package graphql

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen001SourceContract(t *testing.T) {
    source, err := os.ReadFile("context_response.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(c.errors) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
