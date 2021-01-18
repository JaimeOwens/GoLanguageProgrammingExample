package chapter3

import (
	"fmt"
	"os"
)

func Comma() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("	%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
