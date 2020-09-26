package main


import (
    "fmt"
)

func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Printf("The min value is %d", IntMin(2, -2))
}
