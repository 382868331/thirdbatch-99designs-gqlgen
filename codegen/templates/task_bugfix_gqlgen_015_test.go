package templates

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen015SourceContract(t *testing.T) {
    source, err := os.ReadFile("templates.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {") {
        t.Fatalf("expected source contract is missing")
    }
}
