package main

import "fmt"

// no return value of consumer? how to pass the result?
type Consumer = func(int)

func forEach(elems []int, consumer Consumer) {
	for _, elem := range elems {
		consumer(elem)
	}
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	forEach(numbers, func(elem int) {
		sum += elem
	})

	// compare to normal adding iteration
	//for _, elem := range numbers {
	//	sum += elem
	//}
	fmt.Println(sum)
}
