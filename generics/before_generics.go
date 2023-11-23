package generics

import "fmt"

func SumInts(m map[string]int64) int64 {
	var sum int64

	for _, v := range m {
		sum = sum + v
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum = sum + v
	}
	return sum
}

// Test Method to validate the functions before generics
func Test_before_generics() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}
	sumInts := SumInts(ints)

	fmt.Printf("sum of %v is %v\n", ints, sumInts)

	floats := map[string]float64{
		"first":  1.3,
		"second": 2.2,
	}
	sumFloats := SumFloats(floats)

	fmt.Printf("sum of %v is %v\n", floats, sumFloats)
}
