package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestNextToken(t *testing.T) {
	input := `
{
    "foo": "bar",
    "number": 1,
    "height": 0.3,
    "is_real": false,
    "is_good": true, 
    "some_null": null,
    "cars": ["BMW", "Audi", "Benz"]
}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// {
		{token.LCURBRACKET, "{"},
		// "foo": "bar",
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.COMMA, ","},
		// "number": 1,
		{token.STRING, "number"},
		{token.COLON, ":"},
		{token.NUMBER, "1"},
		{token.COMMA, ","},
		// "height": 0.3
		{token.STRING, "height"},
		{token.COLON, ":"},
		{token.NUMBER, "0.3"},
		{token.COMMA, ","},
		// "is_real": false,
		{token.STRING, "is_real"},
		{token.COLON, ":"},
		{token.NUMBER, "false"},
		{token.COMMA, ","},
		// "is_real": false,
		{token.STRING, "is_real"},
		{token.COLON, ":"},
		{token.BOOLEAN, "false"},
		{token.COMMA, ","},
		// "is_good": true,
		{token.STRING, "is_good"},
		{token.COLON, ":"},
		{token.BOOLEAN, "true"},
		{token.COMMA, ","},
		// "some_null": null,
		{token.STRING, "some_null"},
		{token.COLON, ":"},
		{token.NUMBER, "null"},
		{token.COMMA, ","},
		//  "cars": ["BMW", "Audi", "Benz"]
		{token.STRING, "cars"},
		{token.COLON, ":"},
		{token.LSQBRACKET, "["},
		{token.STRING, "BMW"},
		{token.COMMA, ","},
		{token.STRING, "Audi"},
		{token.COMMA, ","},
		{token.STRING, "Benz"},
		{token.RSQBRACKET, "]"},
		{token.RCURBRACKET, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
