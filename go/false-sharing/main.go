package main

import (
	"fmt"
	"sync"
)

func workerThread_1(veryBigSlice []int, index, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := index; i < len(veryBigSlice); i = i + n {
		veryBigSlice[i] = veryBigSlice[i] + 1
	}
}

func false_sharing(in []int) {
	workerCount := 4
	wg := &sync.WaitGroup{}
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go workerThread_1(in, i, workerCount, wg)
	}
	wg.Wait()
}

func workerThread_2(veryBigSlice []int, index, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	cacheLineSize := 64
	for i := index; i < len(veryBigSlice); i = i + n*cacheLineSize/8 {
		for j := 0; j < cacheLineSize && i+j < len(veryBigSlice); j++ {
			veryBigSlice[i+j] = veryBigSlice[i+j] + 1
		}
	}
}

func not_false_sharing(in []int) {
	workerCount := 4
	wg := &sync.WaitGroup{}
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go workerThread_2(in, i, workerCount, wg)
	}
	wg.Wait()
}

func main() {
	veryBigSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	not_false_sharing(veryBigSlice)
	//false_sharing(veryBigSlice)
	fmt.Println("result ", veryBigSlice)
}
