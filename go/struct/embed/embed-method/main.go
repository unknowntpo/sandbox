package main

import "fmt"

type Out struct {
	in In
}

type In struct{}

func (i *In) Hello() {
	fmt.Println("Hello from in")
}

func main() {
	var out Out
	out.Hello()

	//pOut := &out
	//pOut.Hello()
}
