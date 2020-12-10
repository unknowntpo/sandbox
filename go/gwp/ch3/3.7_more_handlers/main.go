package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (f *HelloHandler) S(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (f *WorldHandler) S(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}
func main() {

	// init a handler
	hello := HelloHandler{}
	world := WorldHandler{}

	// Config a server
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	// Register the corresponding handler
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}
