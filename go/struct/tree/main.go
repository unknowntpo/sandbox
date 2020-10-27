package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Insert(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = Insert(t.left, v)
	} else if v > t.value {
		t.right = Insert(t.right, v)
	}
	return t
}

func replaceNode(t, replacement, parent *tree) *tree {
	// check t is at which position of child node of parent
	if t == nil {
		return t
	}
	if t == parent.left {
		parent.left = replacement
		return parent
	} else if t == parent.right {
		parent.right = replacement
		return parent
	}
	return t
}
func findMax(t, parent *tree) (*tree, *tree) {
	if t.right == nil {
		return t, parent
	} else {
		return findMax(t.right, t)
	}
}

// Not familiar of return value handling
func Delete(t, parent *tree, v int) *tree {
	// Search target
	if t == nil {
		return t
	}
	if v < t.value {
		return Delete(t.left, t, v)
	} else if v > t.value {
		return Delete(t.right, t, v)
	}

	// Found target
	// Determine the position of target
	if t.left == nil && t.right == nil {
		// replace parent's child node to nil
		parent = replaceNode(t, nil, parent)
	}
	if t.left == nil {
		parent = replaceNode(t, t.right, parent)
	}
	if t.right == nil {
		parent = replaceNode(t, t.left, parent)
	}
	// t is internal node
	// find max node at t's left subtree
	// and replace t with replacement
	// delete replacement
	replacement, replParent := findMax(t.left, t)
	t.value = replacement.value
	return Delete(replacement, replParent, replacement.value)
}
func Traverse(t *tree) {
	if t == nil {
		return
	}
	// In-order traverse
	Traverse(t.left)
	fmt.Printf("%d ", t.value)
	Traverse(t.right)
	return
}
func test_tree() {
	var r *tree
	r = Insert(r, 5)
	r = Insert(r, 3)
	r = Insert(r, 6)
	r = Insert(r, 4)
	r = Insert(r, 2)

	fmt.Println("In-order traverse ...")
	Traverse(r)
	fmt.Println()
	fmt.Println("Delete node 5")
	r = Delete(r, nil, 5)
	Traverse(r)
}
func main() {
	test_tree()
}
