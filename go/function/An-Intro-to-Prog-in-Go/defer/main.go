package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
func cleanup() {
	fmt.Println("Clean up befor exit ...")
}
func main() {
	defer cleanup()
	first()
	second()
}
