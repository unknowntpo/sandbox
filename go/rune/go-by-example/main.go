package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringToSliceOfRune()
	runeArrayToString()
}

// Transfer string to slice of Rune
func stringToSliceOfRune() {
	s := "ab£"
	r := []rune(s)
	fmt.Println("Input string: ", s)
	fmt.Printf("CodePoint:%U\nType of slice of Rune: %s\n", r, reflect.TypeOf(r))
	fmt.Println("------------")
}

// Transfer rune array to string
func runeArrayToString() {
	a := []rune{'a', 'b', '£'}
	s := string(a)
	fmt.Println("String: ", s)
}
