// Specify the index and value pair without order
// Ref: gopl p.169

package main

import "fmt"


type Currency int
const (
    USD Currency = iota  // how to define constant ?
    EUR
    GBP
    RMB
)

func main() {
    // initialize value with order
    //symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
    // initialize value without order 
    symbol := [...]string{EUR: "€", USD: "$",RMB: "¥", GBP: "£"}

    fmt.Println(symbol)
    fmt.Printf("type of symbol: %T", symbol)
}

