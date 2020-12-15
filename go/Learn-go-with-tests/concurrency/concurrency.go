package concurrency

import (
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// Copy the value of url to avoid access variable 'url' at the same time
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}
	time.Sleep(2 * time.Second)
	return results
}
