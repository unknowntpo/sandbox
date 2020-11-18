package wc

import (
	"strings"
)

type wordCounter int

func (wc *wordCounter) Write(p []byte) (int, error) {
	//split the string using strings.Split, and counter the len out new slice
	words := strings.Split(string(p), " ")
	*wc += wordCounter(len(words))
	return len(words), nil
}
