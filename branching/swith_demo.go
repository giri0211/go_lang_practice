package branching

import (
	"fmt"
)

func Switch_test() {
	x := 5
	switch x {
	case 1:
		fmt.Println("x is 5 from swith")

	case 2 + 3, 2*x + 3:
		fmt.Println(2+3, "case matched to x", x)

	default:
		fmt.Println(x, "default case ")

	}
}

func Logical_switch_test() {

	switch x := 5; true {
	case x < 5:
		fmt.Println("x is 5 from swith")

	case x > 5:
		fmt.Println(2+3, "case matched to x", x)

	default:
		fmt.Println(x, "x is 5 ")

	}
}
