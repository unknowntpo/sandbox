package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

var (
	host = "127.0.0.1"
	port = "8080"
)

func main() {
	parseArgs()

	mux := newRoute()

	n := newLog(mux)

	// Set the parameters for a HTTP server
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: n,
	}

	// Run the server.
	log.Println(fmt.Sprintf("Run the web server at %s:%s", host, port))
	log.Fatal(server.ListenAndServe())
}

func parseArgs() {
	// Consume first argument, which is always the program name.
	args := os.Args[1:]
	for {
		if len(args) < 2 {
			break
		} else if args[0] == "-h" || args[0] == "--host" {
			host = args[1]

			args = args[2:]
		} else if args[0] == "-p" || args[0] == "--port" {
			port = args[1]

			args = args[2:]
		} else {
			log.Fatalln(fmt.Sprintf("Unknown parameter: %s", args[0]))
		}
	}

}
func newRoute() *httprouter.Router {
	// Set a new HTTP request multiplexer
	mux := httprouter.New()

	// Listen to root path
	mux.GET("/", index)

	// Custom 404 page
	mux.NotFound = http.HandlerFunc(notFound)

	// Custom 500 page
	mux.PanicHandler = errorHandler

	return mux
}

func newLog(mux *httprouter.Router) *negroni.Negroni {
	// Set the logger of the server.
	n := negroni.Classic()
	n.UseHandler(mux)
	return n
}
func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Hello World")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Page Not Found")
}

func errorHandler(w http.ResponseWriter, r *http.Request, p interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal Server Error")
}
