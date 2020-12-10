package main

import "fmt"

func main() {
	m := make(map[string]int)

	// modify the value of given key
	m["Answer"] = 42
	fmt.Println(m["Answer"])

	m["Answer"] = 48
	fmt.Println(m["Answer"])

	// Delete an element
	delete(m, "Answer")
	fmt.Println("The value", m["Answer"])

	// Check if key is in map
	elem, ok := m["Answer"]
	fmt.Println("The value: ", elem, "Present?", ok)

}
