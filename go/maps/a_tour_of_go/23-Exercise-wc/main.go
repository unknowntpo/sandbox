package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount
// take an input string s
// return a map
// key: word
// value: amount of occurance of words
// e.g. "I am learning Go!" ->  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
func WordCount(s string) map[string]int {
	// split the string s
	a := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	return m
}
func main() {
	// Use wc package to verify the result
	wc.Test(WordCount)
}
