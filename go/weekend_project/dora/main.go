package main

import (
	"fmt"
	"github.com/bradford-hamilton/dora/pkg/dora"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var exampleJSON = `{ "string": "a neat string", "bool": true, "PI": 3.14159 }`

	c, err := dora.NewFromString(exampleJSON)
	check(err)

	str, err := c.GetString("$.string")
	check(err)

	boolean, err := c.GetBool("$.bool")
	check(err)

	float, err := c.GetFloat64("$.PI")
	check(err)

	fmt.Println(str)     // a neat string
	fmt.Println(boolean) // true
	fmt.Println(float)   // 3.14159
}
