package config

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen005SourceContract(t *testing.T) {
    source, err := os.ReadFile("binder.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "} else if underlying := basicUnderlying(t); def.IsLeafType() && underlying != nil && underlying.Kind() == types.String {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "} else if underlying := basicUnderlying(t); def.IsLeafType() && underlying == nil && underlying.Kind() == types.String {") {
        t.Fatalf("mutated source contract is still present")
    }
}
