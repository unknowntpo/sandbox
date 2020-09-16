// Zero an array using pointer

package main

import "fmt"

func main() {
    c := [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println(c)
    zero(&c)
    //zeroByEmptyArray(&c)
    fmt.Println(c)
}

// zero the contents in array one by one
func zero(pc *[8]byte) {
    for i := range(pc) {
        pc[i] = 0
    }
}

// Assign array literal to the array which pointed to by pc
func zeroByEmptyArray(pc *[8]byte) {
    *pc = [8]byte{}
}
