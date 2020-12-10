package main

import (
	"fmt"
	"strconv"
)

func printInt2String01(num int) string {
	return fmt.Sprintf("%d", num)
}

func printInt2String02(num int64) string {
	return strconv.FormatInt(num, 10)
}
func printInt2String03(num int) string {
	return strconv.Itoa(num)
}
