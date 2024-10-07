package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() { // main project

	var num int
	var num8 int8
	var num16 int16
	var num32 int32
	var num64 int64

	fmt.Printf("num %d byte \n", unsafe.Sizeof(num))
	fmt.Printf("num8 %d byte \n", unsafe.Sizeof(num8))
	fmt.Printf("num16 %d byte \n", unsafe.Sizeof(num16))
	fmt.Printf("num32 %d byte \n", unsafe.Sizeof(num32))
	fmt.Printf("num64  %d byte \n", unsafe.Sizeof(num64))

	var a = bits.UintSize

	fmt.Println(a) // in system

	var flot32 float32 = 100000000000000000000000000000000.200000000000000000000000000000000000000000000000010000000000000000000000000000
	var flot64 float64 = 100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.2320000000000000000000000

	fmt.Printf("flot32 %v  , Size: %d \n", flot32, flot32)  // 3.4 * 10^ -38
	fmt.Printf("flot64 %v , Size: %d  \n ", flot64, flot64) // 1.7 * 10^ -308

}
