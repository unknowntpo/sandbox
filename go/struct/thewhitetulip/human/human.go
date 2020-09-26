// Follow the tutorial from https://thewhitetulip.gitbook.io/bo/02.1introductiongo/02.4struct

package main

import (
	"fmt"
)

type Product struct {
	name          string
	itemID        int
	cost          float32
	isAvailable   bool
	inventoryLeft int
}

/*
Product
|__ goBook

*/
func main() {
    var bookShelf = []Product {
        Product
    }{
        Product{"goBook", 14400, 43.2, true, 32},
        Product{"pythonBook", 14330, 43.1, true, 1},
    }
/*
    var test []struct {
        x, y int
        want int
    }{
        {1, 2, 3},
        {1, 2, 3},
    }
*/
    for _, book := range bookShelf {
        fmt.Printf("book: %+v\n", bookShelf)
    }
}
