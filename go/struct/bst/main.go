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
			n.Right = &Node{Value: value, Data: data}
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
	default:
		return n.Right.Find(s)
	}
}

// Find the maximum element in a tree. Its value replaces the
// value of the to-be-deleted node.
// Return values: the node itself and its parent node
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// replaceNode replaces the parent's child pointer to n
// with a pointer to the replacement node, parent must no be nil
func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode not allowed on a nil Node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

// Delete remove an element from the tree, it's an error to try deleting an
// element that does not exist.
// In order to delete node properly, Delete needs to know the node's parent node,
// parent must not be nil.
func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("cannot delete nil node")
	}

	// Search the target node
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		// Found the target node to be deleted

		// if target is leaf node
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		// If the target node has only one child, replace the target node with its child
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		// If the target node has two child node (target is internal node),
		// find the maximum element in the left subtree, and replace the node's value, data with
		// the replacement's value, data
		replacement, replParent := n.Left.findMax(n)

		n.Value = replacement.Value
		n.Data = replacement.Data

		// remove the replacement node
		return replacement.Delete(replacement.Value, replParent)
	}
}

// From go by example: https://golangbyexample.com/binary-search-tree-in-go/
func (n *Node) inorder() {
	n.inorderRec(n)
}

func (n *Node) inorderRec(node *Node) {
	if node != nil {
		n.inorderRec(node.Left)
		fmt.Println(node.Value)
		fmt.Println(node.Data)
		n.inorderRec(node.Right)
	}
}
func main() {
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}
	root := &Node{}
	for i := 0; i < len(values); i++ {
		err := root.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
	fmt.Print("Sorted values: | ")
	root.inorder()

}
