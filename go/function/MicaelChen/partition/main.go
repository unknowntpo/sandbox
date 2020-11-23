package main

type Rule func(int) bool

func Partition(original []int, rule Rule) ([]int, []int) {
	fit := make([]int, 0)
	notfit := make([]int, 0)

	for _, elem := range original {
		if rule(elem) {
			fit = append(fit, elem)
		} else {
			notfit = append(notfit, elem)
		}
	}
	return fit, notfit
}
