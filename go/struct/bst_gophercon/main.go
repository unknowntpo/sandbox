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

// Tree level Search() method
// Display the result of node level Search method
func (t *Tree) Search(key int) {
	if t.root == nil {
		fmt.Println("Can't search key in nil tree")
		return
	}
	found := t.root.Search(key)
	if found {
		fmt.Println(key, "is found!")
	} else {
		fmt.Println(key, "does not found!")
	}
	return
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

// Tree level insert method
func (t *Tree) Insert(key int) {
	if t.root == nil {
		t.root = &Node{key, nil, nil}
		return
	}
	t.root.insert(key)
	return
}

// Node level insert method
// Because we can't apply (*Node).insert() on a nil node
// so if n is nil, simply return
func (n *Node) insert(key int) {
	if n == nil {
		return
	}
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key, Left: nil, Right: nil}
			return
		}
		n.Left.insert(key)
		return
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key, Left: nil, Right: nil}
			return
		}
		n.Right.insert(key)
		return
	}
	// key == n.Key, node already exist, no need to do anything
}

func (t *Tree) Remove(key int) {
	if t == nil {
		// Can't apply Remove method on a nil tree
		return
	}
	if t.root == nil {
		// Can't apply Remove method on a nil tree
		return
	}
	t.root = t.root.remove(key)
	return
}

// Remove method search for target to remove,
// and return the new node for replacement of the node we just removed
func (n *Node) remove(key int) *Node {
	// Search for target
	if n == nil {
		return n
	}
	if key < n.Key {
		n.Left = n.Left.remove(key)
	} else if key > n.Key {
		n.Right = n.Right.remove(key)
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
		minKey := n.Right.min()

		n.Key = minKey
		// Remove the minKey
		n.Right = n.Right.remove(minKey)
	}
	return n
}

// Return the key of minimum node at root n
func (n *Node) min() int {
	if n.Left == nil {
		return n.Key
	} else {
		return n.Left.min()
	}
}

// Tree level Show() method
func (t *Tree) Show() {
	if t.root != nil {
		t.root.showInOrder()
		fmt.Println()
		return
	}
}

// Node level showInOrder() method
// Show the tree using in-order traversal
// TODO: How to use function value for pre, in, post order traversal?
func (n *Node) showInOrder() {
	if n != nil {
		n.Left.showInOrder()
		fmt.Printf("%d ", n.Key)
		n.Right.showInOrder()
	}
}
func main() {
	// TODO: Write a test for every operation
	// Define our nil root tree
	t := &Tree{nil}
	keys := []int{5, 3, 7, 2, 4, 6, 8}
	for _, key := range keys {
		t.Insert(key)
	}

	// Show the tree
	t.Show()

	// Search the key that exist in tree
	t.Search(3)
	t.Search(4)
	t.Search(10)

	// Delte the key 5
	fmt.Println("rm 5")
	t.Remove(5)
	t.Show()
	fmt.Println("rm 3")
	t.Remove(3)
	t.Show()
	fmt.Println("rm 4")
	t.Remove(4)
	for _, key := range keys {
		fmt.Println("rm ", key)
		t.Remove(key)
		t.Show()
	}
}
