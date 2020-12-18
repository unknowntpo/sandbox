package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	_, _ = http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	_, _ = http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	} else {
		return b
	}

	return
}
