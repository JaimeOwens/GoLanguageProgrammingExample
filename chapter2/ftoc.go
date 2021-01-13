package chapter2

import "fmt"

func F2C() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g。F or %g。C\n", freezingF, fTOC(freezingF))
	fmt.Printf("%g。F or %g。C\n", boilingF, fTOC(boilingF))
}

func fTOC(f float64) float64 {
	return (f - 32) * 5 / 9
}
