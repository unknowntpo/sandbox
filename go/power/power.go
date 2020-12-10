// This program assume base and power >= 0
package power

import "fmt"

func powerRec(base, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return base
	}
	res := powerRec(base, power/2)
	res *= res
	if power&1 == 1 {
		res *= base
	}
	return res
}
func powerIter(base, power int) int {
	res := 1
	for ; power > 0; base *= base {
		if power&1 == 1 {
			res *= base
		}
		power >>= 1
	}
	return res
}
func main() {
	a := 2
	b := 3
	//res := powerIter(a, b)
	res := powerRec(a, b)
	fmt.Println(res)
}
