package singlefile

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGqlgen016SourceContract(t *testing.T) {
    source, err := os.ReadFile("resolver.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "func (r *subscriptionResolver) Updated(ctx context.Context) (<-chan string, error) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "func (r *subscriptionResolver) Updated(ctx context.Context) (<=-chan string, error) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
