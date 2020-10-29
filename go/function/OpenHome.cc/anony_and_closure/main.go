package main

import "fmt"

// Define the type of handling function of filter
type Rule func(int) bool

// filter filter the given slice with handling predicate
func filter(origin []int, rule Rule) []int {
	filtered := []int{}
	for _, elem := range origin {
		// if the elem passed the check of rule(), we preserve it
		if rule(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Input slice: ", data)

	// Pass in data and the handling function to filter
	// Filter the slice by elements less than 5
	fmt.Println("Filter by number that is less than 5")
	fmt.Println(filter(data, func(elem int) bool {
		return elem < 5
	}))

	// Filter the slice by even number
	fmt.Println("Filter by even number")
	fmt.Println(filter(data, func(elem int) bool {
		return 0 == elem&1
	}))

	fmt.Println("Filter by odd number")
	fmt.Println(filter(data, func(elem int) bool {
		return 1 == elem&1
	}))

}
