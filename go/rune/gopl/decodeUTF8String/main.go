package main

import (
	"fmt"
	"unicode/utf8"
)

func CountLengthOfRuneString(s string) {
	fmt.Println(s)
	fmt.Println("number of bytes:", len(s))
	fmt.Println("number of runes:", utf8.RuneCountInString(s))
}
func DecodeUsingRangeLoop(s string) {
	fmt.Println("\nDecodeUsingRangeLoop:")
	fmt.Printf("i\tr\tr\n")
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
func main() {
	s := "Hello, 世界"
	CountLengthOfRuneString(s)
	DecodeUsingRangeLoop(s)
}
