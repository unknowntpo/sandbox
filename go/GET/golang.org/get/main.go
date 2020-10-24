// This is an example from https://golang.org/pkg/net/http/#example_Get
// Sending get request using golang

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./get <URL>")
		os.Exit(1)
	}
	res, err := http.Get(os.Args[1])
	check(err)

	// Read response Header
	fmt.Println("Header of response:")
	fmt.Println(res.Header)

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	check(err)
	res.Body.Close()
	/*
		// Print header of response
		fmt.Println("Header of response:")
		fmt.Printf("%s\n", header)
	*/
	fmt.Println("--------------------------------------------------")

	// Print body of response
	fmt.Println("Body of response:")
	fmt.Printf("%s\n", body)
}
