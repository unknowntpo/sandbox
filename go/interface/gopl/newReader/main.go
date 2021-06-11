package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	//"strings"
)

type Reader struct {
	s string // store the data to be read.
	i int64  // point to the first unread byte.
}

// NewReader convert a string s to Reader.
func NewReader(s string) *Reader {
	return &Reader{s: s, i: 0}
}

// Reader reads up to len(p) bytes from Reader r to buffer p.
// Returns total bytes it read, and error if occured.
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}

	n = copy(p, r.s[r.i:])
	r.i += int64(n)

	return
}

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
	// Implement my own string.NewReader function
	doc, err := html.Parse(NewReader(s))
	//doc, err := html.Parse(strings.NewReader(s))
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
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// What does %*s do?
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
