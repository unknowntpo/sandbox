package main

import "fmt"

type Trie struct {
	root *Node
}

type alphabet int32

type Node struct {
	isEnd    bool
	char     alphabet
	children map[alphabet]*Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{isEnd: true, children: nil}}
}
func (t *Trie) Insert(s string) {
	if t.root.isEnd == true {
		// the root is not the end now
		t.root.isEnd = false
		// insert char
		t.root.insert(s)
	}
	// trie exist
	// search alphabet
}
func (n *Node) insert(s string) {
	// if c is exist in n.children
	// no need to append children
	// if not
	// append children
	hasChar := false
	for _, c := range n.children {
		// if character exist
		// recursively call insert to that alphabet
		if c == a.char {
			a.insert(c)
		}
		hasChar = true
	}
	// if character not exist
	if hasChar == false {
		// check if n.children exist
		// if not, make the map and insert character
		if n.children == nil {
			n.children = map[alphabet]*Node{
				c: &Node{isEnd: true, children: nil},
			}
		}
	}
}

func (t *Trie) Search()

func main() {
	var t *Trie
	t = NewTrie()
	t.Insert("and")
	t.Insert("ant")
	t.Insert("ace")
	//t.Show()
	//found := t.Search("age")
	//fmt.Println(found)
}
