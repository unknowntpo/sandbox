package main

import "fmt"

func main() {
	passwds := map[string]int{"caterpillar": 123456}
	// test existance of key "monica"
	if _, exist := passwds["monica"]; exist {
		fmt.Println("key exist")
	} else {
		fmt.Println("key not exist")
	}

}
