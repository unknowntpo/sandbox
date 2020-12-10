package mysort

//Reference: [Go Slices and Merge Sort](https://golanging.blogspot.com/2013/04/go-slices-and-merge-sort.html)
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

// merge sort
// input: an integer slice
// output: an sorted integer slice
func mergeSort(in []int) []int {
	if in == nil {
		return nil
	} else if len(in) <= 1 {
		return in
	}
	// split
	mid := len(in) / 2
	left := mergeSort(in[:mid])
	right := mergeSort(in[mid:])
	// merge
	return merge(left, right)
}

// Merge two sorted slice
func merge(left, right []int) []int {
	if left == nil && right == nil {
		return []int(nil)
	}
	out := make([]int, 0, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}
	if i < len(left) {
		out = append(out, left[i:]...)
	}
	if j < len(right) {
		out = append(out, right[j:]...)
	}
	return out
}

func main() {
	in := []int{5, 4, 3, 2, 1}
	fmt.Println(bubbleSort(in))
}
