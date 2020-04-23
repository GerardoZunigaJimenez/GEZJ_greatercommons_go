package main

import (
	"GEZJ_greatercommons_go/28_NinjaLevel_13/28_03/mymath"
	"fmt"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println("Slice to be Computed:\n\t", v, "\nResult:\n\t", mymath.CenteredAvg(v), "\n\n")
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	f := []int{5, 3, 2, 4, 1}
	e := [][]int{a, b, c, d, f}
	return e
}
