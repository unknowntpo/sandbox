package main

import (
	"log"
	"net/http"
)

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Any code here will execute on the way down the chain.
		log.Println("middleware1 down")
		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
		log.Println("middleware1 up")
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Any code here will execute on the way down the chain.
		log.Println("middleware2 down")
		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
		log.Println("middleware2 up")
	})
}

func main() {
	mux := http.NewServeMux()
	srv := http.Server{
		Addr:    ":4000",
		Handler: middleware1(middleware2(mux)),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
