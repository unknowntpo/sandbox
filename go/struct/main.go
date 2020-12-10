package main

import (
	"fmt"
	"treesort"
)

func main() {
	t := &tree{}
	t = add(t, 3)
	t = add(t, 2)
	t = add(t, 4)
	t = add(t, 1)
        t = add(t, 5)

        // Store sorted element in the slice
        var values []int
        values = appendValues(values, t)
        fmt.Println(values)
}
