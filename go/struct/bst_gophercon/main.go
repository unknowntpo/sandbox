package main

import "fmt"

const indent = 4
const format = "--<"

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
	if t == nil {
		// TODO: Can't apply Insert method on nil pointer receiver, use error to handle it
		return
	}
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
	if t == nil {
		// TODO: Can't apply Insert method on nil pointer receiver, use error to handle it
		return
	}
	if t.root == nil {
		t.root = &Node{key, nil, nil}
		return
	}
	t.root.insertIter(key)
	//t.root.insert(key)
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

// Implement insert using iteration
func (n *Node) insertIter(key int) {
	for p := n; p != nil; {
		if key < p.Key {
			if p.Left == nil {
				p.Left = &Node{Key: key, Left: nil, Right: nil}
				return
			}
			p = p.Left
		} else if key > p.Key {
			if p.Right == nil {
				p.Right = &Node{Key: key, Left: nil, Right: nil}
				return
			}
			p = p.Right
		}
		// if key == n.Key, do nothing
	}
	return
}
func (t *Tree) Remove(key int) {
	if t == nil || t.root == nil {
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

// Define the function signature of function that passed in to (*Tree).Show()
type ShowFunc func(int)

// Tree level Show() method
func (t *Tree) Show(f ShowFunc) {
	if t == nil {
		// TODO: Can't apply Insert method on nil pointer receiver, use error to handle it
		return
	}
	if t.root != nil {
		fmt.Println("Iterative Pre order traversal:")
		t.root.showPreOrderIter(f)
		//t.root.showInOrder(f)
		fmt.Println()
		return
	}
}

// Node level showInOrder() method
// Show the tree using in-order traversal
// TODO: How to use function value for pre, in, post order traversal?
func (n *Node) showInOrder(f ShowFunc) {
	if n != nil {
		n.Left.showInOrder(f)
		f(n.Key)
		n.Right.showInOrder(f)
	}
}

func (n *Node) showPreOrderIter(f ShowFunc) {
	stack := make([]*Node, 0)
	// Push root to stack
	stack = append(stack, n)
	for len(stack) != 0 {
		// node points to top of stack
		node := stack[len(stack)-1]

		// pop the element at top of stack
		stack = stack[:len(stack)-1]

		f(node.Key)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

//Print CLI readable format of tree
// Ref: [Go Data Structure: Binary Search Tree](https://flaviocopes.com/golang-data-structure-binary-search-tree/)
func (t *Tree) String() {
	if t == nil {
		return
	}

	// Init the depth
	depth = 0
	fmt.Println("--------------------------")
	if t.root != nil {
		stringify(t.root, func(key int) {
			fmt.Printf("%*s%d%s\n", depth*indent, "", key, format)
		})
	}
	fmt.Println("--------------------------")

	// Reset the depth
	depth = 0
}

var depth int

// Internal function to print the tree structure
func stringify(n *Node, f ShowFunc) {
	if n == nil {
		depth++
		fmt.Printf("%*sX\n", depth*indent, "")
		depth--
		return
	} else {
		depth++
		stringify(n.Left, f)
		f(n.Key)
		stringify(n.Right, f)
		depth--
	}
}
func main() {
	// TODO: Write a test for every operation
	// Define our nil root tree
	// Why this line cause panic? -> not handle nil method pointer receiver
	//var t *Tree
	t := &Tree{nil}
	keys := []int{5, 3, 7, 2, 4, 6, 8}
	for _, key := range keys {
		t.Insert(key)
	}

	// Show the tree
	fmt.Println("Tree after Insert", keys, ":")
	t.Show(func(key int) {
		fmt.Printf("%d ", key)
	})

	// Search the key that exist in tree
	t.Search(3)
	t.Search(4)
	t.Search(10)

	// Show CLI readable tree structure
	fmt.Println("Show CLI readable tree structure...")
	t.String()

	// Remove all keys
	fmt.Println("Remove all keys...")
	for _, key := range keys {
		fmt.Println("rm ", key)
		t.Remove(key)
		t.Show(func(key int) {
			fmt.Printf("%d ", key)
		})
	}

}
