package main

type Lexer struct {
	input        string
	position     int // points to current read position
	readPosition int // points to next read position
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return &Lexer{input}
}

func (l *Lexer) readChar() {
	if theEnd(readPosition, len(input)) {
		l.ch = 0
	} else {
		l.ch = input[readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// theEnd report whether we are at read end.
// e.g. input: {"foo": "bar"}, if position is at '}', we are at the end
func theEnd(pos int, length int) bool {
	return pos == length
}

// TODO: Implement nextToken
func (l *Lexer) nextToken() token.Token {

}
