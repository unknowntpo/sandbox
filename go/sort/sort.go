package sort

import "fmt"

// bubbleSort
// input: an integer slice
// output: an sorted integer slice
func bubbleSort(in []int) []int {
	if len(in) == 0 {
		return in
	}
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
	return in
}
func main() {
	in := []int{5, 4, 3, 2, 1}
	fmt.Println(bubbleSort(in))
}
