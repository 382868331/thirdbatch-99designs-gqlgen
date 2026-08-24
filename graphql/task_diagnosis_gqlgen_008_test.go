package graphql

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGqlgen008SourceContract(t *testing.T) {
    source, err := os.ReadFile("uint.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if _, err := strconv.ParseUint(v[1:], 10, 64); err == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && _, err := strconv.ParseUint(v[1:], 10, 64); err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
