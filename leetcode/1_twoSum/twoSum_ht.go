package main

import (
	"fmt"
	"github.com/unknowntpo/hashmap"
)

func main() {
	m := &hashmap.Map{}
	bit := 10
	m.Init(bit)
	nums := []int{2, 7, 11, 15}

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m.Add(v, i)
	}

	for i, _ := range nums {
		v, ok := m.Get(nums[i]) // expect 0, true
		if ok {
			fmt.Printf("nums[%d]: %d\n", i, v)
		} else {
			fmt.Printf("nums[%d] doesn't exist!", i)
		}
	}

	m.String()
}
