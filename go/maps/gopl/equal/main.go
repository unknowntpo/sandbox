package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	// Loop through all element in map and check equal
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}
func main() {
	x := map[string]int{"a": 1, "b": 2}
	y := map[string]int{"a": 1, "b": 2, "c": 3}

	isequal := equal(x, y)
	fmt.Println("Are x, y equal?", isequal)
}
