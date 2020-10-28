package main

import "fmt"

func main() {
	fmt.Println("Panic is comming (❁°͈▵°͈)")
	defer func() {
		if r := recover(); r != nil {
			// Display the panic message
			fmt.Println("What's happended?", r)
			fmt.Println("Recover from panic finally ...")
		} else {
			// No panic, nothing to recover from
			fmt.Println("Nothing happened, don't worry!")
		}
	}()
	panic("Panic")

}
