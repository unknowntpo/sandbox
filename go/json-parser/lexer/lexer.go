package main

import (
	"fmt"

	"../token"
)

type Lexer struct {
	input        string
	position     int // points to current read position
	readPosition int // points to next read position
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if theEnd(l.readPosition, len(l.input)) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// theEnd report whether we are at read end.
// e.g. input: {"foo": "bar"}, if position is at '}', we are at the end
func theEnd(pos int, length int) bool {
	return pos >= length
}

// TODO: Implement NextToken
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	fmt.Printf("we got a %q\n", l.ch)
	l.skipWhitespace()
	switch l.ch {
	case '{':
		return newToken(token.LCURBRACKET, '{')
	default:
		fmt.Printf("we got a %q\n!", l.ch)
	}
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
