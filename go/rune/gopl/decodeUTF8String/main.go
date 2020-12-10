package main

import (
	"fmt"
	"unicode/utf8"
)

func CountLengthOfRuneString(s string) {
	fmt.Println(s)
	fmt.Println("number of bytes:", len(s))

	fmt.Println("Count number of runes:")
	fmt.Println("\tUsing utf-8 package:", utf8.RuneCountInString(s))
	RangeCountRunes(s)

}
func RangeCountRunes(s string) {
	fmt.Printf("\tUsing range loop: ")
	n := 0
	for range s {
		n++
	}
	fmt.Println(n)

}
func DecodeUsingRangeLoop(s string) {
	fmt.Println("\nDecodeUsingRangeLoop:")
	fmt.Printf("i\tr\tr\n")
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// Decode utf-8 encoding string using utf-8 package
	fmt.Println("\nDecode utf-8 encoding string using utf-8 package")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
func main() {
	s := "Hello, 世界"
	CountLengthOfRuneString(s)
	DecodeUsingRangeLoop(s)
}
