package branching

import (
	"fmt"
)

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()

	//create a panic by diving by 0
	result := dividend / divisor
	// fmt.Println("after division statement")
	return result
}

func Create_panic_manually(dividend, divisor int) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()

	fmt.Println("before panic")

	panic("hello panic")

}

func Divide_test() {
	dividend, divisor := 10, 5
	fmt.Printf("%v dividend is divided by %v divisor and result is %v", dividend, divisor, divide(dividend, divisor))
	fmt.Println()
	dividend, divisor = 10, 0
	fmt.Printf("%v dividend is divided by %v divisor and result is %v", dividend, divisor, divide(dividend, divisor))
	fmt.Println()
}
