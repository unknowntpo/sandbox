package main

import (
	"testing"
)

/*
Package, go doc support
*/

// Helper function for filling the tree by given test slice
func (tree *Tree) fillTree(keys []int) {
	for _, key := range keys {
		tree.Insert(key)
	}
}

// Helper function for checking if the input slices are the same
// Return true if a, b are the same.
// Also return true if a, b are all nil slices
// Return false if a, b are not the same.
// Ref: https://flaviocopes.com/golang-data-structure-binary-search-tree/
func isSameSlice(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestShow(t *testing.T) {
	var tests = []struct {
		keys []int
		want []int
	}{
		{[]int{5, 3, 7, 2, 4, 6, 8}, []int{2, 3, 4, 5, 6, 7, 8}},
	}

	for _, test := range tests {
		var tree = &Tree{nil}
		tree.fillTree(test.keys)

		// Traverse the tree and append the elements into slice
		var result []int
		tree.Show(func(key int) {
			result = append(result, key)
		})
		// Check if the result are the same slice as we wanted
		if !isSameSlice(result, test.want) {
			t.Errorf("In order traverse incorrect, get %d, want %d", result, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		name   string
		keys   []int
		target int
		want   []int
	}{
		{"Remove root node", []int{5, 3, 7, 2, 4, 6, 8}, 5, []int{2, 3, 4, 6, 7, 8}},
		{"Apply Remove on nil tree", []int(nil), 5, []int(nil)},
		{"Remove leaf node", []int{5, 3, 7, 2, 4, 6, 8}, 2, []int{3, 4, 5, 6, 7, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var tree = &Tree{nil}
			tree.fillTree(test.keys)

			// Traverse the tree and append the elements into slice
			var result []int
			tree.Remove(test.target)
			tree.Show(func(key int) {
				result = append(result, key)
			})
			// Check if the result are the same slice as we wanted
			if !isSameSlice(result, test.want) {
				t.Errorf("Remove method failed, get %d, want %d", result, test.want)
			}
		})
	}

}
