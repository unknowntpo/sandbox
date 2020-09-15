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
	// initialize it
	var goBook Product
	// initialization of goBook individually
	goBook.name, goBook.itemID, goBook.isAvailable, goBook.inventoryLeft = "Webapps in Go", 10025, true, 25
	// init of pythonBook by format "field:value" without order
	pythonBook := Product{itemID: 10026, name: "Learn Python", inventoryLeft: 0, isAvailable: false}
	// init rubyBook in order
	rubyBook := Product{"LearnRuby", 10043, 100, true, 12}

	if goBook.isAvailable {
		fmt.Printf("%d copies of %s are available\n",
			goBook.inventoryLeft, goBook.name)
	}

	if rubyBook.isAvailable {
		fmt.Printf("%d copies of %s are available\n",
			rubyBook.inventoryLeft, rubyBook.name)
	}

	if pythonBook.isAvailable {
		fmt.Printf("%d copies of %s are available\n",
			pythonBook.inventoryLeft, pythonBook.name)
	}

}
