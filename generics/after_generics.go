package generics

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func SumOfNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum = sum + v
	}
	return sum
}

// Test Method to validate the generic function.
func Test_after_generics() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}
	sumInts := SumOfNumbers(ints)

	fmt.Printf("sum of %v is %v\n", ints, sumInts)

	floats := map[string]float64{
		"first":  1.3,
		"second": 2.2,
	}
	sumFloats := SumOfNumbers(floats)

	fmt.Printf("sum of %v is %v\n", floats, sumFloats)
}
