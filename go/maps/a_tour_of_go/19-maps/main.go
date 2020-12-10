package main

import "fmt"

// declare a map: key: string; value: struct Vertex: Lat, Long
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
