package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (n *Node) Search(key int) bool {
	// key not found
	if n == nil {
		return false
	}
	if key < n.Key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}
	// key == n.Key, found target
	return true
}

func (t *Tree) Insert(key int) {
	if t.root == nil {
		t.root = &Node{key, nil, nil}
		return
	}
	t.root.Insert(key)
	return
}

// Node level insert method
// Because we can't apply (*Node).Insert() on a nil node
// so if n is nil, simply return
func (n *Node) Insert(key int) {
	if n == nil {
		return
	}
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key, Left: nil, Right: nil}
			return
		}
		n.Left.Insert(key)
		return
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key, Left: nil, Right: nil}
			return
		}
		n.Right.Insert(key)
		return
	}
	// key == n.Key, node already exist, no need to do anything
}

// Delete method search for target to delete,
// and return the new node for replacement of the node we just deleted
func (n *Node) Delete(key int) *Node {
	// Search for target
	if key < n.Key {
		n.Left = n.Left.Delete(key)
	} else if key > n.Key {
		n.Right = n.Right.Delete(key)
	} else {
		// found the target
		// Determine the delete method
		if n.Left == nil {
			// We will reach here if n is either a leaf node, or n's Left child is nil
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// If we reach here, means n is an internal node
		// Find the Min number of node at n's right subtree
		// to replace the position of n
		min := n.Right.Min()

		n.Key = min
		n.Right = n.Right.Delete(min)
	}
	return n
}

// Return the key of minimum node at root n
func (n *Node) Min() int {
	if n.Right == nil {
		return n.Key
	} else {
		return n.Left.Min()
	}
}

// Show the tree given different method of traversing
// TODO: How to use function value for pre, in, post order traversal?
func (n *Node) Show() {
	if n != nil {
		n.Left.Show()
		fmt.Printf("%d ", n.Key)
		n.Right.Show()

	}
}
func main() {
	// Define our nil root tree
	t := &Tree{nil}
	t.Insert(5)
	t.Insert(3)
	t.Insert(7)
	t.Insert(2)
	t.Insert(4)
	t.Insert(6)
	t.Insert(8)
	t.root.Show()
}
