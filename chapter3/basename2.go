package chapter3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Basename2() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename2(input.Text()))
	}
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
