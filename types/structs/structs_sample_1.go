package structs

import "fmt"

type person struct {
	name string
	age  int
}

// func printPassedValue(p *person) {
// 	fmt.Printf("struct passed by address %v\n", *p)
// }
func printPassedByValue(p person) {
	p.name = "byvalue"
}

func printPassedByRef(p *person) {
	p.name = "byRef"
}

func Test_structs() {

	person1 := person{name: "girish", age: 35}
	fmt.Printf("struct create by property names %v\n ", person1)

	x := person{name: "girish"}
	x.age = 10
	fmt.Printf("struct assinged only name %v\n", x)
	fmt.Printf("struct assinged only name %v\n", &x)

	printPassedByValue(x)
	fmt.Printf("after printPassedByValue %v\n", x)
	printPassedByRef(&x)
	fmt.Printf("after printPassedByValue %v\n", x)

	y := x
	y.name = "changed by value"
	fmt.Printf("y changed by value %v\n", y)
	fmt.Printf("x after y changed by value %v\n", x)
}
