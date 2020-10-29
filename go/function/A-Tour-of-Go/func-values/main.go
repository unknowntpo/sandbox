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
func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func main() {

	// hypot return the hypotenuse of given triangle
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// perim return the perimeter of given triangle
	perim := func(x, y float64) float64 {
		return x + y + hypot(x, y)
	}

	// compute the hypotenuse of triangle with side 3, 4
	fmt.Println("Hypotenuse of triangle with side 3, 4: ", compute(hypot, 3, 4))
	// compute x to the power of y
	fmt.Println("3 to the power of 4: ", compute(math.Pow, 3, 4))
	// compute the perimeter of given triangle
	fmt.Println("Perimeter of triangle with side 5, 6: ", compute(perim, 5, 6))
}
