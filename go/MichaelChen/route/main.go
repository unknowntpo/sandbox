package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Params:", p)
	fmt.Fprintf(w, "Hello, %s!", p.ByName("name"))
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Hello World")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	// Why use WriteHeader?
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Page not Found handled by me")
}

func errorHandler(w http.ResponseWriter, r *http.Request, p interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal Server Error handled by me")
}
func main() {
	// Create a new HTTP request ultiplexer
	mux := httprouter.New()

	// Listen to root path
	mux.GET("/", index)

	// custom 404 page
	// why using HandlerFunc?
	mux.NotFound = http.HandlerFunc(notFound)

	// Custom 500 page
	mux.PanicHandler = errorHandler

	// visit a site == GET method?
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
