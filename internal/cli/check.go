package cli

import (
	"fmt"
	"io"

	"github.com/MustafaRady20/opsbot/internal/checker"
	"github.com/spf13/cobra"
)

// defaultTargets is used until T04 loads targets from the config file.
var defaultTargets = []string{
	"https://www.google.com",
	"https://doesnotexist.invalid",
	"https://github.com",
}

func newCheckCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "Check every target once and print the results",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCheck(cmd.OutOrStdout(), defaultTargets)
		},
	}
}

// runCheck checks each URL, prints one line per result to w, and returns
// an error if any target is down, so the process exits with code 1.
func runCheck(w io.Writer, urls []string) error {
	down := 0
	for _, u := range urls {
		r := checker.CheckHTTP(u)
		fmt.Fprintln(w, r)
		if !r.OK {
			down++
		}
	}
	if down > 0 {
		return fmt.Errorf("%d of %d targets down", down, len(urls))
	}
	return nil
}
