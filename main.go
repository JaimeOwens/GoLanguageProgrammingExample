package main

import (
	"./chapter2"
	"fmt"
)

func main() {
	fmt.Println(chapter2.CToF(chapter2.AbsoluteZeroC))
	fmt.Println(chapter2.CToF(chapter2.FreezingC))
	fmt.Println(chapter2.CToF(chapter2.BoilingC))
}
