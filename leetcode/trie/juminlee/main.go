package main

import "fmt"

type Trie struct {
	root *Node
}

type Node struct {
	c     string
	m     alphabet
	isEnd bool
}

type alphabet map[string]*Node

func (t *Trie) Insert(s string) {
	cur := t.root
	for i, v := range s {
		if v_in_cur_m() {
			cur = t.m[v]
			if t.isEnd {
				n := &Node{c: v, m: make(alphabet), isEnd: true}
				cur.m[v] = n
				cur = cur.m[v]
				continue
			}
		} else {
			n := &Node{c: v, m: make(alphabet), isEnd: true}
			cur.m[v] = n
		}
		// if v in cur.m

		// // if move cur point to t.m["v"]
		// else

	}

}

func main() {
	fmt.Println("vim-go")
}
