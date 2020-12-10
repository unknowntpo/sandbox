package main

import "fmt"

type Rule func(int) int

// mapping input to output by given rule
func mapping(origin []int, rule Rule) []int {
	var out []int

	for _, elem := range origin {
		out = append(out, rule(elem))
	}
	return out
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("origin slice: ", s)
	// Multiply by two
	fmt.Println("Mul by 2:", mapping(s, func(elem int) int {
		return elem * 2
	}))

	// Increment by one
	fmt.Println("Increment by 1:", mapping(s, func(elem int) int {
		return elem + 1
	}))

}
