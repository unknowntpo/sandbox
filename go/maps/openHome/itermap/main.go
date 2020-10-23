package main

import (
	"fmt"
	"sort"
)

func getkeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getvalues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
func main() {
	passwords := map[string]int{
		"caterpillar": 123456,
		"monica":      54321,
	}
	// iterate the key and value in map
	for name, value := range passwords {
		fmt.Println("Key: ", name, "Value: ", value)
	}
	keys := getkeys(passwords)
	fmt.Println(keys)

	values := getvalues(passwords)
	fmt.Println(values)

	// use sort to sort the map
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, passwords[key])
	}

}
