package config

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen007SourceContract(t *testing.T) {
    source, err := os.ReadFile("config.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if r := recover(); r != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
