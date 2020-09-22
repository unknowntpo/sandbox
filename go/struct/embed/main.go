package main

import (
	"fmt"
)

type Point struct{ X, Y int }

type Circle struct {
	Point  // doesn't need data type ?
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel

	// Declare a wheel with: Circle: X: 8 Y: 8, Radius: 5 spokes: 20
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{ // Specified field by index: value, just like Point
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
                Spokes: 20,
	}

	fmt.Printf("%+v\n", w)

	w.X = 42

	fmt.Printf("%#v\n", w)
}
