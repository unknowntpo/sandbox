package main

import "fmt"

func reverse(s string) string {

	in := []byte(s)
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return string(in)
}
func main() {
	in := "abcdefg"
	fmt.Println("input", in)
	out := reverse(in)
	fmt.Println("output: ", out)
}
