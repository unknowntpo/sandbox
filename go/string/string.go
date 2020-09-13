/*
    Experiment: Can we modify the data of string?
    Ref: 
        - [My Gopl note](https://hackmd.io/@unknowntpo/Golang/%2FhIMTN172RfurPW7QHW7foQ#Goal-echogo-%E4%BA%86%E8%A7%A3%E5%85%A7%E5%AE%B9-p32)
        - [GOPL Ch3.5 Strings p.139]
*/
package main

import (
    "fmt"
)

func main() {
    s := "Hello"
    s[1] = 'h'
    fmt.Println(s)
}
