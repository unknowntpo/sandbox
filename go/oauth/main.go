package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	CLIENT_ID     = os.Getenv("GITHUB_OAUTH_CLIENT_ID")
	CLIENT_SECRET = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", loginHandler)

	port := 4444

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("oauth example app listen on :%d", port)
	log.Fatal(srv.ListenAndServe())
}
