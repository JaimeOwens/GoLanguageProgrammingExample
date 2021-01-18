package chapter3

import (
	"bufio"
	"fmt"
	"os"
)

func Basename1() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename1(input.Text()))
	}
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
