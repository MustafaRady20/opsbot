package cli

import (
	"bytes"
	"strings"
	"testing"
)

// execute runs the command tree with args and returns what it printed.
func execute(t *testing.T, args ...string) (string, error) {
	t.Helper()
	var out bytes.Buffer
	root := NewRootCmd("v1.2.3")
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs(args)
	err := root.Execute()
	return out.String(), err
}

func TestVersion(t *testing.T) {
	got, err := execute(t, "version")
	if err != nil {
		t.Fatalf("version returned error: %v", err)
	}
	if want := "opsbot v1.2.3\n"; got != want {
		t.Errorf("version printed %q, want %q", got, want)
	}
}

func TestHelpListsCommands(t *testing.T) {
	got, err := execute(t, "--help")
	if err != nil {
		t.Fatalf("--help returned error: %v", err)
	}
	for _, want := range []string{"check", "version", "--config"} {
		if !strings.Contains(got, want) {
			t.Errorf("--help output is missing %q:\n%s", want, got)
		}
	}
}

func TestUnknownCommandFails(t *testing.T) {
	if _, err := execute(t, "nope"); err == nil {
		t.Error("unknown command returned no error")
	}
}
