package todo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen013SourceContract(t *testing.T) {
    source, err := os.ReadFile("todo.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if todo.ID == id {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if todo.ID != id {") {
        t.Fatalf("mutated source contract is still present")
    }
}
