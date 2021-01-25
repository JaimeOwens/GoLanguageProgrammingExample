package chapter13

import (
	"fmt"
	"unsafe"
)

func UnsafePtr() {
	var x struct {
		a bool
		b int16
		c []int
	}

	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b)
}
