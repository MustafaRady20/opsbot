package checker

import (
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
