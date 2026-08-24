package modelgen

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen019SourceContract(t *testing.T) {
    source, err := os.ReadFile("models.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "} else if !interfaceFieldTypeIsPointer && structFieldTypeIsPointer {") {
        t.Fatalf("expected source contract is missing")
    }
}
