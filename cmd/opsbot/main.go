package main

import (
	"fmt"
	"time"

	"github.com/MustafaRady20/opsbot/internal/checker"
)

func main() {
	fmt.Println("opsbot dev")

	urls := []string{
		"https://www.google.com",
		"https://doesnotexist.invalid",
		"https://github.com",
	}

	for _, u := range urls {
		r := checker.CheckHTTP(u)
		latency := r.Latency.Round(time.Millisecond)
		switch {
		case r.Err != nil:
			fmt.Printf("DOWN %s  error: %v\n", r.Target, r.Err)
		case !r.OK:
			fmt.Printf("DOWN %s  status %d  %v\n", r.Target, r.Status, latency)
		default:
			fmt.Printf("UP   %s  status %d  %v\n", r.Target, r.Status, latency)
		}
	}
}
