package main

import (
	"fmt"
	//	"github.com/unknowntpo/hashmap"
)

/*
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := &hashmap.Map{}
	bit := 10
	m.Init(bit)
	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m.Add(v, i)
	}

	m.String()
	for _, num := range nums {
		wantKey := target - num
		fmt.Println(wantKey)
		// look up map to get key index
		wantKeyIdx, ok := m.Get(wantKey)
		if ok {
			res = appendIdxUnique(res, wantKeyIdx)

		}
	}

	return res
}
*/
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)

	m := make(map[int]int)

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m[v] = i
	}

	for _, num := range nums {
		wantKey := target - num
		// Make sure the value of wantKey exist and wantKey not equal to num itself, because it doesn't match the meaning to twoSum, e.g. [3, 2, 4] -> if key == 3, then wantKey will be 3, too.
		if wantKey == num {
			continue
		}
		// look up map to get key index
		wantKeyIdx, ok := m[wantKey]
		if ok {
			res = appendIdxUnique(res, wantKeyIdx)
		}
	}

	return res
}

func appendIdxUnique(idxSlice []int, idx int) []int {

	for i := range idxSlice {
		if idx == idxSlice[i] {
			return idxSlice
		}
	}

	// if num not in nums
	idxSlice = append(idxSlice, idx)
	return idxSlice
}
func main() {
	target := 6
	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 4}
	res := twoSum(nums, target)
	fmt.Println(res)
}
