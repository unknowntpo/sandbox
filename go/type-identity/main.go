package main

import (
	"fmt"
	"net/http"
)

func main() {
	mockHandlerFunc(http.MethodGet, "/", homeHandler)
}

func mockHandlerFunc(method, path string, handler http.HandlerFunc) {
	var w http.ResponseWriter
	var r *http.Request

	fmt.Printf("%T\n", handler)

	handler.ServeHTTP(w, r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
}
