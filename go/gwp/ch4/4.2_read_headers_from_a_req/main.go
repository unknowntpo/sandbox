package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Headers of response:")
	for k, _ := range r.Header {
		fmt.Fprintln(w, k)
		if v, ok := r.Header[k]; ok {
			fmt.Fprintln(w, v)
		}
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//Data Type of server?
	fmt.Printf("%T\n", server)
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
