package chapter7

import (
	"flag"
	"fmt"
)

//!+
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func TemoFlag() {
	flag.Parse()
	fmt.Println(*temp)
}
