// Zero an array using pointer

package main

import "fmt"

func main() {
    c := [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println(c)
    zero(&c)
    fmt.Println(c)
}

func zero(ptr *[8]byte) {
    // If you only need the first item in the range (the key or index), drop the second:
    for i := range ptr {
        ptr[i] = 0
    }
}

