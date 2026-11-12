// Package checker checks whether network targets are healthy.
package checker

import (
	"net/http"
	"time"
)

// Result is the outcome of a single check against one target.
type Result struct {
	Target  string
	OK      bool
	Status  int
	Latency time.Duration
	Err     error
}

// CheckHTTP sends a GET request to url and reports whether it answered
// with a 2xx status. Err is set when the request itself failed, for example
// on a DNS error or a refused connection; Status is then zero.
func CheckHTTP(url string) Result {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return Result{
			Target:  url,
			Latency: time.Since(start),
			Err:     err,
		}
	}
	defer resp.Body.Close()

	return Result{
		Target:  url,
		OK:      resp.StatusCode >= 200 && resp.StatusCode < 300,
		Status:  resp.StatusCode,
		Latency: time.Since(start),
	}
}
