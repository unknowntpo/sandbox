package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	// GWP's version of reading body
	/*
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	*/
	// my version of read reading the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	fmt.Fprintln(w, string(body))
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
