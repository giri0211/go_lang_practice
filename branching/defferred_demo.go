package branching

import "fmt"

func Deffered_function_test() {

	fmt.Println("statement 1")
	defer fmt.Println("defer statement 1")
	fmt.Println("statement 2")
	defer fmt.Println("defer statemenet 2")
	// defer is last in first out ordering
}
