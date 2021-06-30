package main

import "fmt"

type Token int

const (
	TokenReserved = Token(iota)
	TokenNumber
	TokenWhiteSpace
	TokenInvalid
)

// checkType checks type of token of c, and return Token type.
func checkType(c byte) Token {
	if c == '+' || c == '-' {
		return TokenReserved
	}

	if c == ' ' {
		return TokenWhiteSpace
	}

	if c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
		return TokenNumber
	}

	return TokenInvalid
}

// tokenizer fetch characters in s and return slice of byte contains
// tokens.
func tokenizer(s string) []byte {
	out := []byte{}
	for i := 0; i < len(s); i++ {
		switch checkType(s[i]) {
		case TokenWhiteSpace:
		case TokenNumber:
			out = append(out, s[i])
		case TokenReserved:
			out = append(out, s[i])
		}
	}
	return out
}

func showTokens(in []byte) {
	for i, v := range in {
		fmt.Printf("idx[%d]: -> %s\n", i, string(v))
	}
}
func main() {
	s := "1 + 2 + 3"
	t := tokenizer(s)
	showTokens(t)
}
