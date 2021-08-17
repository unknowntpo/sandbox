package main

import "fmt"

type Calculator struct {
	acc float64
}

const (
	OP_ADD = 1 << iota
	OP_SUB
	OP_MUL
)

func (c *Calculator) Do(op int, v float64) float64 {
	switch op {
	case OP_ADD:
		c.acc += v
	case OP_SUB:
		c.acc -= v
	case OP_MUL:
		c.acc *= v
	default:
		panic("unhandled operation")
	}
	return c.acc
}

func main() {
	var c Calculator
	fmt.Println(c.Do(OP_ADD, 100)) // 100
	fmt.Println(c.Do(OP_SUB, 50))  // 50
	fmt.Println(c.Do(OP_MUL, 2))   // 100
}
