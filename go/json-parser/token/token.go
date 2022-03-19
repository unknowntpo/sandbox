package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	OBJECT  = "OBJECT"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"
	NUMBER  = "NUMBER"
	ARRAY   = "ARRAY"
	NULL    = "NULL"

	COLON       = ":"
	COMMA       = ","
	LCURBRACKET = "{"
	RCURBRACKET = "}"
	LSQBRACKET  = "["
	RSQBRACKET  = "]"
)

/*
EBNF

object = "{", key-value-pair, "}";

key-value-pair = string, ":", (string | object | array | number | null);

string = { "A" | "B" | "C" | "D" | "E" | "F" | "G"
       | "H" | "I" | "J" | "K" | "L" | "M" | "N"
       | "O" | "P" | "Q" | "R" | "S" | "T" | "U"
       | "V" | "W" | "X" | "Y" | "Z" | "a" | "b"
       | "c" | "d" | "e" | "f" | "g" | "h" | "i"
       | "j" | "k" | "l" | "m" | "n" | "o" | "p"
       | "q" | "r" | "s" | "t" | "u" | "v" | "w"
       | "x" | "y" | "z" } ;

array = "[" , object {",", object}, "]";
*/
/*
{
  "level": "fatal",
  "key": {
      array: [
	...
      ]
  }
  "error": "dial tcp [::1]:5432: connect: connection refused",
  "time": "2022-03-12T07:11:46+08:00"
}
*/
