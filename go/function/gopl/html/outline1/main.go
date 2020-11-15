package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// Fetch the url in or.Args[1]
// And return the resp as type of *http.Response
func fetch() *http.Response {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	return resp
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		// What is n.Data?
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
func main() {
	// handle cmd arg.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./outline1 <URL>\n")
		os.Exit(1)
	}
	resp := fetch()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline1: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Display the html DOM tree
	outline(nil, doc)
}
