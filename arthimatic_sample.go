package main

import "fmt"

func arthimatic_test() {
	a, b := 5, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println(float32(a) / float32(b))

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
}
