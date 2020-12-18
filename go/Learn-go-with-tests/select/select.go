package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	} else {
		return b
	}

	return
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
