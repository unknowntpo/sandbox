package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world!")
}
func main() {
	// set up the server
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// set the handler
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	// start the server
	server.ListenAndServe()
}
