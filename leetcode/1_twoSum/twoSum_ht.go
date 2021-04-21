package main

import (
//"fmt"
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
func twoSumGoMap(nums []int, target int) []int {
	res := make([]int, 0)

	m := make(map[int]int)
	// for edge case [3 3] target = 6, which means that refcnt[3] == 2
	refcnt := make(map[int]int)

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m[v] = i
		refcnt[v]++
	}

	for i, num := range nums {
		// take number 'num' from nums
		refcnt[num]--
		wantKey := target - num
		// look up map to get key index
		wantKeyIdx, ok := m[wantKey]
		if ok {
			// if no more key "wantKey" to take
			if refcnt[wantKey] > 0 {
				refcnt[wantKey]--
				res = appendIdxUnique(res, wantKeyIdx)
				res = appendIdxUnique(res, i)
			}
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
}
