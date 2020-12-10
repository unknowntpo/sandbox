package main

import "fmt"

// Print the vertical triangle with given number of layer n
func printVertTriangle(n int) {
	spaceNum := n - 1
	asteriskNum := n - spaceNum
	for i := n; i > 0; i, spaceNum, asteriskNum = i-1, spaceNum-1, asteriskNum+2 {
		for j := 0; j < spaceNum; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j < asteriskNum; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < spaceNum; j++ {
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}
func main() {
	n := 5
	for i := 0; i < n; i++ {
		printVertTriangle(i)
		fmt.Println()
	}
}
