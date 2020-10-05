// Implement a stack using slice

package main

import "fmt"

// Remove the element in slice[idx] and shift all elements after idx-th element
// 1 position left to fill the gap
func remove(slice []int, idx int) []int {
	copy(slice[idx:], slice[idx+1:])
	return slice[:len(slice)-1]
}

// Remove the element specified by index by
// replacing the gap by the last element
func removeWithoutOrder(slice []int, idx int) []int {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func create(value ...int) []int {
	return value
}

func top(s []int) {
	fmt.Printf("%d\n", s[len(s)-1])
}

// TODO: Use pointer to prevent returning back s
func push(s []int, value int) []int {
	s = append(s, value)
	return s
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func show(s []int) {
	fmt.Printf("stack: %v\n", s)
}
func main() {

	fmt.Println("creating new stack...")
	stack := create(5, 6, 7, 8)
	//stack := []int{5, 6, 7, 8, 9}
	show(stack)
	// Find the top of stack
	fmt.Println("Top of stack...")
	top(stack)

	// Push 1 element into the stack
	fmt.Println("Pushing...")
	stack = push(stack, 9)
	show(stack)

	// Pop 1 element out of the stack
	fmt.Println("Popping...")
	stack = pop(stack)
	show(stack)

	// Remove 1 element out of the stack
	fmt.Println("Removing...")
	show(remove(stack, 2)) // [5, 6, 8, 9]
	//show(removeWithoutOrder(stack, 2))
}
