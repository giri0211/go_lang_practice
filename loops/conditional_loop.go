package loops

import (
	"fmt"
	"strings"
)

// infinite loop
func Infinite_loop_test() {
	fmt.Println("Infinite_loop_test")
	fmt.Println(strings.Repeat("_", 10))
	i := 99
	for {
		fmt.Println(i)
		i++
		break
	}
}

// condition based loop
func Conditional_loop_test() {
	fmt.Println("Conditional_loop_test")
	fmt.Println(strings.Repeat("_", 10))
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// counter based loop
func Counter_based_loop_test() {
	fmt.Println("Counter_based_loop_test")
	fmt.Println(strings.Repeat("_", 10))
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
