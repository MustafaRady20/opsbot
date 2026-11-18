package main

import (
	"os"

	"github.com/MustafaRady20/opsbot/internal/cli"
)

// version is overwritten at build time with
// -ldflags "-X main.version=...". It stays "dev" for go run.
var version = "dev"

func main() {
	if err := cli.NewRootCmd(version).Execute(); err != nil {
		os.Exit(1)
	}
}
