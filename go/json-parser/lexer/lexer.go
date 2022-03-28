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

	l.skipWhitespace()
	fmt.Printf("we got a %q at input[%d], readPos: %d\n", l.ch, l.position, l.readPosition)
	switch l.ch {
	case '{':
		l.readChar()
		return newToken(token.LCURBRACKET, "{")
	case ':':
		l.readChar()
		return newToken(token.COLON, ":")
	case ',':
		l.readChar()
		return newToken(token.COMMA, ",")
	default:
		switch {
		case l.ch == '"':
			s := l.getUntilNextPair('"')
			fmt.Println(s)
			l.readChar()
			return newToken(token.STRING, string(s))
		case isDigit(l.ch):
			num := l.readNumber()
			fmt.Println("num from readNumber", num)
			// TODO: Implement readNumber
		}
		fmt.Printf("we got a %q\n!", l.ch)

	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	// e.g.
	// 5.3
	// 4.6
	// 0.1
	// 1
	// 12343
	// .1 // valid or invalid
	var out []byte
	out = append(out, l.input[l.position])

	// number of dots we encountered
	once := false
	for {
		fmt.Println("cur l.ch", string(l.ch))
		// how to deal with invalid number ?
		// we ignore it for now
		if isDot(l.ch) && once == true {
			// we got two dots, we just return current output
			return string(out)
		} else {
			once = true
			goto APPEND

		}
		if !isDigit(l.ch) {
			return string(out)
		}

	APPEND:
		// it's a digit
		out = append(out, l.ch)
		l.readChar()
	}
	return ""
}

func isDot(c byte) bool {
	return c == '.'
}

// 0.1     out: 0.1

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

//
/*

e.g.
123456
"hello"
*/
func (l *Lexer) getUntilNextPair(c byte) string {
	start := l.readPosition
	var idx int

	// advance 1 position, or we won't get into the loop because l.ch == c
	l.readChar()
	for idx = start; l.ch != c; idx++ {
		fmt.Println("idx", idx)

		l.readChar()
	}
	return l.input[start:idx]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
