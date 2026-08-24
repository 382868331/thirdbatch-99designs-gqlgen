package transport

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGqlgen018SourceContract(t *testing.T) {
    source, err := os.ReadFile("websocket_coderws.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil && isCoderNormalClose(err) {") {
        t.Fatalf("expected source contract is missing")
    }
}
