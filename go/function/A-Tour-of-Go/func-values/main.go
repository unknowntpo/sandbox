package main

import (
	"fmt"
	"math"
)

// function value in go == function pointer in c ?
// compute the property of two values 3, 4 using the handle func passed in
// e.g.
// 	compute34(hypot): compute the hypotenuse of triangle with side 3, 4
//	compute34(math.Pow): compute 3 to the power of 4
func compute34(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	// hypot return the hypotenuse of given triangle
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Hypotenuse of triangle with side 5, 12: ", hypot(5, 12))

	// compute the hypotenuse of triangle with side 3, 4
	fmt.Println("Hypotenuse of triangle with side 3, 4: ", compute34(hypot))
	// compute x to the power of y
	fmt.Println("3 to the power of 4: ", compute34(math.Pow))
}
