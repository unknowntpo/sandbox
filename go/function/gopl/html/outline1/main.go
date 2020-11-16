package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

const s = `
<html>
    <head>
	<title>Example HTML </title>
    </head>
    <body>
	<h1>Hello world</h1>
	<h2>Sub head </h2>
    </body>
</html>`

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
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline1: %v\n", err)
		os.Exit(1)
	}

	// Display the html DOM tree
	outline(nil, doc)
}
