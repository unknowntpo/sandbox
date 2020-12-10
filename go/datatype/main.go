package main

import (
	"fmt"
	"unsafe"
)

const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7 // 如果用 1 << 7 會 overflow ? why -1 << 7 不會？
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func main() {
	// What's the size and data type of constant ?
	fmt.Printf("MaxInt8: type: %T, size: %d\n", MaxInt8, unsafe.Sizeof(MaxInt8))
	fmt.Printf("MinInt8: type: %T, size: %d\n", MinInt8, unsafe.Sizeof(MinInt8))
	fmt.Printf("MaxInt16: type: %T, size: %d\n", MaxInt16, unsafe.Sizeof(MaxInt16))
	fmt.Printf("MinInt16: type: %T, size: %d\n", MinInt16, unsafe.Sizeof(MinInt16))
	fmt.Printf("MaxInt32: type: %T, size: %d\n", MaxInt32, unsafe.Sizeof(MaxInt32))
	fmt.Printf("MinInt32: type: %T, size: %d\n", MinInt32, unsafe.Sizeof(MinInt32))

}
