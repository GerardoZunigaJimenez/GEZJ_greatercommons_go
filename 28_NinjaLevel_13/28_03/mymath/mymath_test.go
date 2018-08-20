package mymath

import (
	"testing"
	"fmt"
)

type test struct {
	slice  []int
	result float64
}

var sTest = []test{
	test{[]int{1, 4, 6, 8, 100}, 6,},
	test{[]int{0, 8, 10, 1000}, 9,},
	test{[]int{9000, 4, 10, 8, 6, 12}, 9,},
	test{[]int{123, 744, 140, 200}, 170,},
}

func TestCenteredAvg(t *testing.T) {
	for i := 0; i < len(sTest); i++ {
		result := CenteredAvg(sTest[i].slice)
		if result != sTest[i].result {
			t.Error("Expected:", sTest[i].result, "Got:", result)
		}
	}
}

func ExampleCenteredAvg() {
	result := CenteredAvg([]int{5, 3, 2, 4, 1})
	fmt.Println(result)
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(sTest); i++ {
			result := CenteredAvg(sTest[i].slice)
			if result != sTest[i].result {
				b.Error("Expected:", sTest[i].result, "Got:", result)
			}
		}
	}
}
