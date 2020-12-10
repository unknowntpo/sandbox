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

type Tree struct {
	Root *Node
}

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

// Wrapper method for insert node in tree
func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}

	// call the node-level insert method
	return t.Root.Insert(value, data)
}

// Node-level insert method
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

// Wrapper method for finding specific node
func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

// Node-level find method
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

// Wrapper method for deleting node
func (t *Tree) Delete(s string) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// What if we don't use fake parent?
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}

	// Why need this checck?
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
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

	if parent == nil {
		return errors.New("parent node must not be nil")
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

// Pre-Order traverse
func (t *Tree) Traverse_Pre(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	f(n)
	t.Traverse_Pre(n.Left, f)
	t.Traverse_Pre(n.Right, f)
}

// In-Order traverse
func (t *Tree) Traverse_In(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse_In(n.Left, f)
	f(n)
	t.Traverse_In(n.Right, f)
}

// Post-Order traverse
func (t *Tree) Traverse_Post(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse_Post(n.Left, f)
	t.Traverse_Post(n.Right, f)
	f(n)
}

// Level-Order traverse
func (t *Tree) Traverse_Level(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	f(n)

	if n.Left != nil {
		f(n.Left)
	}
	if n.Right != nil {
		f(n.Right)
	}

	t.Traverse_Level(n.Left, f)
	t.Traverse_Level(n.Right, f)
}

// Wrapper function for traversing a tree
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	// Pre-Order traverse
	fmt.Println("Pre order traverse:")
	t.Traverse_Pre(n, f)
	fmt.Println()

	// In-Order traverse
	fmt.Println("In order traverse:")
	t.Traverse_In(n, f)
	fmt.Println()

	// Post-Order traverse
	fmt.Println("Post order traverse")
	t.Traverse_Post(n, f)
	fmt.Println()

	// Level-Order traverse
	fmt.Println("Level order traverse")
	t.Traverse_Level(n, f)
	fmt.Println()
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

	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value `", values[i], "':", err)
		}
	}

	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Test find

	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	// Test deletion
	s = "d"
	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Test for single node tree
	fmt.Println("Single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	// Node level operation Demo
	/*
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
	*/
}
