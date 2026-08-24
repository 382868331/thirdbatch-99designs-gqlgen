package starwars

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGqlgen004SourceContract(t *testing.T) {
    source, err := os.ReadFile("resolvers.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return 0, errors.New(\"invalid unit\")") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return 1, errors.New(\"invalid unit\")") {
        t.Fatalf("mutated source contract is still present")
    }
}
