package main

// auto go fmt in vim ?

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("Original")
	fmt.Println(a)

	fmt.Println("First rotate")
	reverse(a[:2])
	fmt.Println(a)

	fmt.Println("Second rotate")
	reverse(a[2:])
	fmt.Println(a)

	fmt.Println("Last rotate")
	reverse(a[:])
	fmt.Println(a)
}
