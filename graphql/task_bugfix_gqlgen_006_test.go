package graphql

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen006SourceContract(t *testing.T) {
    source, err := os.ReadFile("resolve_field.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if r := recover(); r != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
