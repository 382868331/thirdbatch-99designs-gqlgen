package graphql

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGqlgen020SourceContract(t *testing.T) {
    source, err := os.ReadFile("batch.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "groups = make(map[string]*BatchParentGroup, len(prev.groups)+1)") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "groups = make(map[string]*BatchParentGroup, len(prev.groups)- 1)") {
        t.Fatalf("mutated source contract is still present")
    }
}
