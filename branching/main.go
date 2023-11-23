package branching

import "fmt"

func If_statement_test() {
	i := 4
	if i < 5 {
		fmt.Println("i is less than 5")
	}
	fmt.Println("after the if statement")
}

func If_else_statement_test() {

	i := 5
	if i < 5 {
		fmt.Println("i is less than 5")
	} else {
		fmt.Println("else caluse executed ")
	}

	fmt.Println("after the if statement")

}

func If_else_if_statement_test() {

}

func If_else_if_else_statement_test() {

	i := 5
	if i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10 ")
	} else {
		fmt.Println("i is atlest 10 ")
	}

	fmt.Println("after the if statement")

}

func If_initilizer_test() {
	if i := 5; i < 5 {
		fmt.Println("i is less than 5")
	}
}
