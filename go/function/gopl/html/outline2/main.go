package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

const indent = 4
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

func main() {
	outline(s)
}
func outline(s string) error {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return err
	}

	// What does this do?
	forEachNode(doc, startElement, endElement)

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		// pre-order traversal
		pre(n)
	}

	// Traverse through child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	// If post is defined, do post-order traversal
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// What does %*s do?
		fmt.Printf("%*s<%s>\n", depth*indent, "", n.Data)
		depth++
	}
	if n.Type == html.TextNode {
		// Control the indent
		fmt.Printf("%*s", depth*indent, "")
		for _, ele := range n.Data {
			if ele == '\n' {
				fmt.Printf("\\n")
			} else if ele == '\t' {
				fmt.Printf("\\t")
			} else {
				fmt.Print(string(ele))
			}
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// What does %*s do?
		fmt.Printf("%*s</%s>\n", depth*indent, "", n.Data)
	}
	if n.Type == html.TextNode {
		depth--
		fmt.Println()
	}
}
