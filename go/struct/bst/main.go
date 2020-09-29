// This code is the practice of bst
// Ref: [Go by example: Binary Search Tree in Go (Golang)](https://golangbyexample.com/binary-search-tree-in-go/)

package main

import "fmt"

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

func (b *bstnode) reset() {
	b.root = nil
}

func (b *bstnode) insert(value int) {
	b.insertRec(b.root, value)
}

// insertRec: recursively search end of node and add new node on it
func (b *bstnode) insertRec(node *bstnode, value int) *bstnode {
	if b.root == nil {
		b.root = &bstnode{
			value: value,
		}
		return b.root
	}
	if node == nil {
		return &bstnode{value: value}
	}
	if value <= node.value {
		node.left = b.insertRec(node.left, value)
	} else {
		node.right = b.insertRec(node.right, value)
	}
	return node
}

func (b *bstnode) find(value int) error {
	node := b.findRec(b.root, value)
	if node == nil {
		return fmt.Errorf("Value: %d not found in tree", value)
	}
	return nil
}

func (b *bstnode) findRec(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.findRec(node.left, value)
	}
	return b.findRec(node.right, value)

}

// What does in order mean?
func (b *bstnode) inorder() {
	b.inorderRec(b.root)
}

func (b *bstnode) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.value)
		b.inorderRec(node.right)
	}
}

func main() {
	bst := &bstnode{}
	eg := []int{2, 5, 7, -1, -1, 5, 5}
	for _, val := range eg {
		bst.insert(val)
	}
	fmt.Printf("Printing Inorder:\n\b")
	bst.inorder()
	bst.reset()
	eg = []int{4, 5, 7, 6, -1, 99, 5}
	for _, val := range eg {
		bst.insert(val)
	}
	fmt.Printf("\nPrinting Inorder:\n")
	bst.inorder()
	fmt.Printf("\nFinding Values:\n")
	err := bst.find(2)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 2)
	} else {
		fmt.Printf("Value %d Found\n", 2)
	}

	// Find some value
	err = bst.find(6)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 6)
	} else {
		fmt.Printf("Value %d Found\n", 6)
	}
	err = bst.find(5)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 5)
	} else {
		fmt.Printf("Value %d Found\n", 5)
	}
	err = bst.find(1)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", 1)
	} else {
		fmt.Printf("Value %d Found\n", 1)
	}
}
