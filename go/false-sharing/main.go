package main

import (
	"fmt"
	"sync"
	//"unsafe"
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

func workerThread_2(veryBigSlice []int, startIdx, workerCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := startIdx; i < len(veryBigSlice); i = i + elesInALine*workerCount {
		for j := 0; j < elesInALine && i+j < len(veryBigSlice); j++ {
			veryBigSlice[i+j]++
		}
	}
}

func not_false_sharing(in []int) {
	workerCount := 4
	wg := &sync.WaitGroup{}
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go workerThread_2(in, i*elesInALine, workerCount, wg)
	}
	wg.Wait()
}

const cacheLineSize = 64 // bytes
// TODO: Use unsafe.Sizeof to get size of int
//elesInALine := cacheLineSize / *(*int)(unsafe.Pointer(unsafe.Sizeof(int(0))))
const elesInALine = cacheLineSize / 4

func main() {
	veryBigSlice := make([]int, 64)
	for i := 0; i < 64; i++ {
		veryBigSlice[i] = i
	}

	not_false_sharing(veryBigSlice)
	//false_sharing(veryBigSlice)
	fmt.Println("result ", veryBigSlice)
}
