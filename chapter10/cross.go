package chapter

import (
	"fmt"
	"runtime"
)

func Cross() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
