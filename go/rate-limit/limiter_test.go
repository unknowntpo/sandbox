package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExceed(t *testing.T) {
	t.Run("Should be exceed if too many request", func(t *testing.T) {
		// init the limiter with rps = 5, burst = 5
		rps := 5
		burst := 5
		addr := Addr("1.1.1.1")
		l := NewTokenBucketLimiter(rps, burst)
		// hit it 5 times in 1 second
		for i := 0; i < 5; i++ {
			allow := l.Allow(addr)
			assert.True(t, true, "Should be allowed ")
		}
		// The 6th time
		// assert rate limit exceeded
		allow := r.Allow()
		assert.False(t, true, "Should be denied because too many request")
	})
}
