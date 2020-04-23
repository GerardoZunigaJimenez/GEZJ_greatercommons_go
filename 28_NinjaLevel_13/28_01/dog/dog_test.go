package dog

import (
	"fmt"
	"testing"
)

type test struct {
	years  int
	dogAge int
}

var tests = []test{
	test{1, 7},
	test{2, 14},
	test{3, 21},
	test{4, 28},
	test{5, 35},
	test{6, 42},
	test{7, 49},
	test{8, 56},
	test{9, 63},
	test{10, 70},
	test{11, 77},
	test{12, 84},
	test{13, 91},
	test{14, 98},
	test{15, 105},
}

func TestYears(t *testing.T) {
	var result int
	for _, v := range tests {
		result = Years(v.years)
		if result != v.dogAge {
			t.Error("Expedted", v.dogAge, "Got", result)
		}
	}
}

func ExampleYears() {
	dogYears := Years(2)
	fmt.Println(dogYears)
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			result = Years(v.years)
			if result != v.dogAge {
				b.Error("Expedted", v.dogAge, "Got", result)
			}
		}
	}
}

func TestYearsTwo(t *testing.T) {
	var result int
	for _, v := range tests {
		result = YearsTwo(v.years)
		if result != v.dogAge {
			t.Error("Expedted", v.dogAge, "Got", result)
		}
	}
}

func ExampleYearsTwo() {
	dogYears := YearsTwo(2)
	//fmt.Println(dogYears)
	// Output:
	// 14
}

func BenchmarkYearsTwo(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			result = YearsTwo(v.years)
			if result != v.dogAge {
				b.Error("Expedted", v.dogAge, "Got", result)
			}
		}
	}
}
