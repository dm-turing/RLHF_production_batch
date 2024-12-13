package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// Test existing functionality
func TestPerformAction(t *testing.T) {
	// Capture output
	storeStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PerformAction()

	w.Close()
	buf, _ := io.ReadAll(r)

	output := string(buf)
	if !strings.Contains(output, "123") {
		t.Errorf("expected output to contain token, got: %s", output)
	}
	os.Stdout = storeStdout // Restore the original output
}

func TestMyAction_PerformAction(t *testing.T) {
	config := Config{APIToken: "test-token"}
	action := MyAction{config: config}

	result := action.PerformAction()
	expected := "Performing action with token: test-token"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
