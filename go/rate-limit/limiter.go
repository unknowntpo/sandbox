package main

import (
	"sync"
)

type Addr string

type RateLimiter interface {
	Allow(Addr) bool
}

type TokenBucketLimiter struct {
	buckets map[Addr]bucket
	mu      sync.Mutex
}

type bucket struct {
	count int
}

// rps: request per second
// burst: capacity of bucket, if burst = 0, means infinite capacity.
func NewTokenBucketLimiter(rps int, burst int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		buckets: make(map[Addr]bucket),
	}
}

const (
	allow = true
	deny  = false
)

func (t *TokenBucketLimiter) Allow(addr Addr) bool {
	if t.inBuckets(addr) {
		if t.notExceed(addr) {
			addOneToBucketCount(addr)
			// allowed
			return allow
		} else {
			// denied because rate limit exceed.
			return deny
		}
	} else {
		// not in buckets, create one
		// TODO: finish implementation
		createInstanceInBuckets(addr)
		// TODO: finish implementation
		addOneToBucketCount(addr)
		return allow
	}
}

func (t *TokenBucketLimiter) inBuckets(addr Addr) bool {
	_, ok := t.buckets[addr]
	return ok
}

func (t *TokenBucketLimiter) notExceed(addr Addr) bool {
	return b.buckets[addr].count > 0
}
