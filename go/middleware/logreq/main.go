package main

import (
	"log"
	"net/http"
)

type application struct {
}

func (app *application) logReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO]: %s %s from %s", r.Method, r.URL.String(), r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
func demoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I see you"))
}

func (app *application) route() http.Handler {
	route := http.NewServeMux()
	route.HandleFunc("/demo", demoHandler)
	return app.logReq(route)
}

func main() {
	app := &application{}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.route(),
	}

	log.Println("starting server at ", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
