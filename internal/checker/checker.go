package checker

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Target  string
	OK      bool
	Status  int
	Latency time.Duration
	Err     error
}

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

// String formats r as one line, for example "UP   https://example.com  status 200  120ms".
// It makes Result a fmt.Stringer, so fmt.Println(r) prints this line.
func (r Result) String() string {
	latency := r.Latency.Round(time.Millisecond)
	switch {
	case r.Err != nil:
		return fmt.Sprintf("DOWN %s  error: %v", r.Target, r.Err)
	case !r.OK:
		return fmt.Sprintf("DOWN %s  status %d  %v", r.Target, r.Status, latency)
	default:
		return fmt.Sprintf("UP   %s  status %d  %v", r.Target, r.Status, latency)
	}
}
