// This code is the practice of bst
// Ref:
//       - [Applied Go: A binary search tree](https://appliedgo.net/bintree/)
//       - [My HackMD](https://hackmd.io/@unknowntpo/bst-explain)
package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value, data string) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {

	case value == n.Value:
		return nil
	case value < n.Value:
		// go to left node
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		// go to right node
		if n.Right == nil {
			n.right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s string) (string, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	case s > n.Value:
		return n.Right.Find(s)
	}
}
