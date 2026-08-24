package modelgen

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen009SourceContract(t *testing.T) {
    source, err := os.ReadFile("interface_graph.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(schemaType.Interfaces) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
