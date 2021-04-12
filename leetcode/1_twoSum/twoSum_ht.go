package main

import (
	"fmt"
)

type Map struct {
	bit int
	ht  []htHead
}

type htHead struct {
	first *htNode // points to first element of that bucket
}

type htNode struct {
	key, value int
	next       *htNode
}

func (m *Map) Init(bit int) {
	m.bit = bit
	size := 1 << bit
	m.ht = make([]htHead, size)
}

func (h *htHead) appendHtNode(key, value int) {
	node := h.first
	if node == nil {
		node = &htNode{key: key, value: value, next: nil}
		return
	}
	for ; node.next != nil; node = node.next {
		if node.key == key {
			// key exist, no need to append
			return
		}
	}

	// at tail of list, append newNode
	newNode := &htNode{key: key, value: value, next: nil}
	node.next = newNode

}

func (m *Map) Add(key, value int) {
	set := key % (2 << m.bit)
	// access m.ht[set]
	// if m.ht[set] not exist, append m.ht[]
	/*
		for len(m.ht) < set {
			m.ht = append(m.ht, htHead{first: nil})
		}
	*/
	// then append htNode
	m.ht[set].appendHtNode(key, value)
}

func (m *Map) Get(key int) (value int, ok bool) {
	set := key % (2 << m.bit)
	// access m.ht[set]
	// if m.ht[set] not exist, means key not exist in map
	if len(m.ht) < set {
		return 0, false
	}

	for node := m.ht[set].first; node != nil; node = node.next {
		if node.key == key {
			// key exist, return the value and ok
			return node.value, true
		}
	}
	// key not found
	return 0, false
}

// String shows the whole map with format Map[key:value]
func (m *Map) String() {
	// figure out some method to contain all keys
}

func main() {
	m := &Map{}
	bit := 10
	m.Init(bit)
	nums := []int{2, 7, 11, 15}
	for i, v := range nums {
		// add element of nums as key, and the index of element in nums as value
		m.Add(v, i)
	}
	v, ok := m.Get(nums[0]) // expect 0, true
	if ok {
		fmt.Println(nums[0], ":", v)
	} else {
		fmt.Println("not exist !")
	}

}
