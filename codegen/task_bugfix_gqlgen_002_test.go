package codegen

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen002SourceContract(t *testing.T) {
    source, err := os.ReadFile("field.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if found != nil || err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if found == nil || err != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
