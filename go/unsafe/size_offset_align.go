package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(x.a))

	fmt.Println(unsafe.Sizeof(x.b))

	fmt.Println(unsafe.Sizeof(x.c))
	// FIXME: How to get sizeof uintptr ?

	// FIXME: How to get sizeof slice.ptr?
	//fmt.Println(unsafe.Sizeof(unsafe.Pointer(x.a)))

	// Offsetof
	fmt.Println(unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Offsetof(x.c))

}
