package main

import "fmt"

func main() {
	fmt.Println("Panic is comming (❁°͈▵°͈)")
	defer func() {
		str := recover()
		fmt.Println(str)
		fmt.Println("Recover from panic finally ...")
	}()
	panic("Panic")
}
