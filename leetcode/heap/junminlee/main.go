// Ref: Data Structures and Algorithms in Go - Heaps
// https://www.youtube.com/watch?v=3DYIgTC4T1o&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6&ab_channel=JunminLee
package main

import "fmt"

type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap.
// return -1 if can't extract.
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	extracted := h.array[0]
	l := len(h.array) - 1

	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		switch {
		// when left child is the only child
		case l == lastIndex:
			childToCompare = l
		// when left child is larger
		case h.array[l] > h.array[r]:
			childToCompare = l
			// when right child is larger
		default:
			childToCompare = l
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

	for h.array[child(index)] > h.array[index] {
		h.swap(child(index), index)
		index = child(index)
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func child(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)

	for i := 0; i < 10; i++ {
		fmt.Print(m.Extract(), "\t")
		fmt.Println(m)
	}
}
