package hello

import (
	"rsc.io/quote"
	quotev3 "rsc.io/quote/v3"
)

func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return quotev3.Concurrency()
}
