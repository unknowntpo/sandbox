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
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		// Append the extracted link from DOM tree to slice: links
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// What does next sibling do?
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
func main() {
	// handle cmd arg.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./findlinks1 <URL>\n")
		os.Exit(1)
	}
	resp := fetch()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
