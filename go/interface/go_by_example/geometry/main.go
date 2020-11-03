package main

import (
	"fmt"
	"math"
)

// define an interface with area and perimeter method
// take in 2 named method rect and circle
// define a measuring function taking in interface and output the result

type geometry interface {
	perim() float64
	area() float64
}

type rect struct {
	length, width float64
}
type circle struct {
	radius float64
}

func (c *circle) perim() float64 {
	return math.Pi * c.radius * 2
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *rect) perim() float64 {
	return 2 * (r.length + r.width)
}

func (r *rect) area() float64 {
	return r.length * r.width
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Perimeter: ", g.perim())
	fmt.Println("Area: ", g.area())
}
func main() {
	r := &rect{2, 4}
	c := &circle{3.5}
	measure(r)
	measure(c)
}
