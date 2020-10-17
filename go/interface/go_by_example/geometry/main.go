package main

import (
	"fmt"
	"math"
)

// define an interface with area and perimeter method
// take in 2 named method rect and circle
// define a measuring function taking in interface and output the result

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perim() float64 {
	return 2 * r.length + 2 * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("interface: %+v\n", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	r := rect{2, 4}
	c := circle{3.5}
	measure(r)
	measure(c)
}
