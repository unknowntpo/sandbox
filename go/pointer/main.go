// This is a practice of using pointer to pass data by reference
// Ref: https://thewhitetulip.gitbook.io/bo/02.1introductiongo/02.3cntrlstmtfunctions#pass-by-value-and-pointers

package main

import (
    "fmt"
)

func add1(a *int) {
    *a += 1
    return
}

func main() {
    a := 1
    add1(&a)
    fmt.Println("Result of addition", a)
}
